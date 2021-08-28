<template>
  <div class="signup-container">
    <el-form ref="signUpForm" :model="signUpForm" :rules="loginRules" class="signup-form" autocomplete="on" label-position="left">
      <div class="title-container">
        <h3 class="title">Sign Up Form</h3>
      </div>
      <el-form-item prop="user_name">
        <span class="svg-container">
          <svg-icon icon-class="user" />
        </span>
        <el-input
          ref="user_name"
          v-model="signUpForm.user_name"
          placeholder="username <email>"
          name="user_name"
          type="text"
          tabindex="1"
          autocomplete="on"
        />
      </el-form-item>

      <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
        <el-form-item prop="password">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="signUpForm.passwd"
            :type="passwordType"
            placeholder="Password"
            name="password"
            tabindex="2"
            autocomplete="on"
            @keyup.native="checkCapslock"
            @blur="capsTooltip = false"
            @keyup.enter.native="handleLogin"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>
      </el-tooltip>
      <el-tooltip v-model="capsTooltip" content="Caps lock is On" placement="right" manual>
        <el-form-item prop="password">
          <span class="svg-container">
            <svg-icon icon-class="password" />
          </span>
          <el-input
            :key="passwordType"
            ref="password"
            v-model="signUpForm.re_passwd"
            :type="passwordType"
            placeholder="Password"
            name="re_password"
            tabindex="3"
            autocomplete="on"
            @keyup.native="checkCapslock"
            @blur="capsTooltip = false"
            @keyup.enter.native="handleLogin"
          />
          <span class="show-pwd" @click="showPwd">
            <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
          </span>
        </el-form-item>
      </el-tooltip>
      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="handleSignUp">SignUp</el-button>

      <div style="position:relative">
        <div class="tips">
          <span>Username : email</span>
          <span>Password : 6-12char</span>
        </div>
      </div>
    </el-form>
  </div>
</template>

<script>
import { sleep } from '@/utils/auth'
export default {
  name: 'SignUp',
  data () {
    const validateUsername = (rule, value, callback) => {
      console.log('validateUsername', value)
      if (value.length < 4) {
        callback(new Error('the name can not be less than 4 digits '))
      } else {
        callback()
      }
    }
    const validatePassword = (rule, value, callback) => {
      console.log(value)
      if (value.length < 6) {
        callback(new Error('The password can not be less than 6 digits'))
      } else {
        callback()
      }
    }
    return {
      signUpForm: {
        user_name: '',
        passwd: '',
        re_passwd: '',
        email: ''
      },
      loginRules: {
        user_name: [
          { required: true, trigger: 'blur', validator: validateUsername },
          { min: 3, max: 25, trigger: 'blur', message: '长度在 3 到 25 个字符' }
        ],
        passwd: [{ required: true, trigger: 'blur', validator: validatePassword }],
        re_passwd: [{ required: true, trigger: 'blur', validator: validatePassword }]
      },
      passwordType: 'password',
      capsTooltip: false,
      loading: false,
      redirect: undefined,
      otherQuery: {}
    }
  },
  methods: {
    checkCapslock (e) {
      const { key } = e
      this.capsTooltip = key && key.length === 1 && (key >= 'A' && key <= 'Z')
    },
    showPwd () {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleSignUp () {
      this.$refs.signUpForm.validate(valid => {
        if (valid) {
          this.loading = true
          this.$store.dispatch('user/signup', this.signUpForm)
            .then(() => {
              console.log('login success' + this.redirect)
              this.$router.push({ path: this.redirect || '/', query: this.otherQuery })
              this.loading = false
            })
            .catch((err) => {
              console.log('login faild' + err + this.redirect)
              this.loading = false
            })
        } else {
          console.log('error submit!!')
          return false
        }
      })
    }
  }
}
</script>

<style lang="scss" scoped>
$bg:#2d3a4b;
$dark_gray:#889aa4;
$light_gray:#eee;
    .signup-container {
      min-height: 100%;
      width: 100%;
      background-color: $bg;
      overflow: hidden;
      .signup-form {
        position: relative;
        width: 520px;
        max-width: 100%;
        padding: 60px 35px 0;
        margin: 0 auto;
        overflow: hidden;
      }

      .tips {
        font-size: 14px;
        color: #fff;
        margin-bottom: 10px;

        span {
          &:first-of-type {
            margin-right: 16px;
          }
        }
      }

      .svg-container {
        padding: 6px 5px 6px 15px;
        color: $dark_gray;
        vertical-align: middle;
        width: 30px;
        display: inline-block;
      }

      .title-container {
        position: relative;

        .title {
          font-size: 26px;
          color: $light_gray;
          margin: 0px auto 40px auto;
          text-align: center;
          font-weight: bold;
        }
      }

      .show-pwd {
        position: absolute;
        right: 10px;
        top: 7px;
        font-size: 16px;
        color: $dark_gray;
        cursor: pointer;
        user-select: none;
      }

      .thirdparty-button {
        position: absolute;
        right: 0;
        bottom: 6px;
      }

      @media only screen and (max-width: 470px) {
        .thirdparty-button {
          display: none;
        }
      }
    }
</style>
