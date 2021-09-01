<template>
  <div class="signup-container">
    <el-form ref="signUpForm" class="signup-form" autocomplete="on" label-position="left">
      <div class="title-container">
        <h3 class="title">扫码登录</h3>
      </div>
      <el-form-item prop="user_name">
        <div class="demo-image" >
          <div class="block" align="center" style="background-color: #fff;">
            <el-image
                style="width: 200px; height: 200px; margin: 20px"
                :src="imageUrl"
                fit="cover"></el-image>
          </div>
        </div>
      </el-form-item>
      <div class="msg-container">
         <h5 class="title">{{msg}}</h5>
      </div>
      <el-button :loading="loading" type="primary" style="width:100%;margin-bottom:30px;" @click.native.prevent="getQrLogin">刷新</el-button>
    </el-form>
  </div>
</template>

<script>
import { sleep } from '@/utils/auth'
import Command from '@/api/jdseckill/qrLogin'
import {Message} from "element-ui";
export default {
  name: 'JdLogin',
  data () {
    return {
      timerQrTick: '',
      timerQrToken: '',
      wlfstkSmdl:'',
      imageUrl:"",
      code: 0,
      msg:'',
      tick:'',
      loading: false,
      redirect: undefined,
      otherQuery: {}
    }
  },

  mounted() {
    this.timerQrTick = setInterval(this.getQrTick, 2000);
    this.timerQrToken = setInterval(this.getQrToken, 2000);
  },
  methods: {
    getQrLogin () {
      //Message.info('do getQrLogin')
      this.wlfstkSmdl = ''
      const data = {}
      Command.send_show(data).then(response => {
        console.log('data:', response.data)
        //Message.info('imageUrl' + response.data.imgUrl)
        this.wlfstkSmdl = response.data.wlfstkSmdl
        this.imageUrl = response.data.imgUrl
      }).catch(error => {
        console.log('delete error:', error)
        Message.error({ message: error })
      })
    },
    getQrTick () {
      if (this.wlfstkSmdl==''){
        return
      }
      //Message.info('do getTick ' + this.wlfstkSmdl)
      const data = {}
      Command.send_get_tick(data, this.wlfstkSmdl).then(response => {
        //console.log('data:', response.data)
        this.code = response.data.code
        this.msg = response.data.msg
        if (response.data.ticket){
          this.tick = response.data.ticket
          this.wlfstkSmdl = ''
        }
        if (response.data.msg=='参数错误'){
          this.wlfstkSmdl = ''
        }
      }).catch(error => {
        console.log('delete error:', error)
        //Message.error({ message: error })
      })
    },
    getQrToken () {
      if (this.tick==''){
        return
      }
      //Message.info('do get_token ' + this.tick)
      const data = {}
      Command.send_get_token(data, this.tick).then(response => {
        console.log('data:', response)
        if (response.token){
          // this.thor = response.data.thor
          // this.unick = response.data.unick
          this.tick= ''
          Message.info('登录成功 ')
          console.log('登录成功:', response)
          this.$store.dispatch('user/setToken', response.token)
          this.$router.push('/')
        }
        this.msg = this.unick = response.data.msg
        this.tick= ''
      }).catch(error => {
        console.log('delete error:', error)
        //Message.error({ message: error })
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
      .msg-container {
        position: relative;

        .title {
          font-size: 12px;
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
