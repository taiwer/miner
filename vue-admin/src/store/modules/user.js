import { login, logout, signUp, getInfo, refreshToken } from '@/api/user'
import { setToken, getToken, removeToken, setTokenExpire, getTokenExpire, removeTokenExpire } from '@/utils/auth'
import router, { resetRouter } from '@/router'
import { sleep } from '../../utils/auth'

const state = {
  name: '',
  avatar: '',
  introduction: '',
  roles: [],
  token: '',
  tokenExpire: ''
}

const mutations = {
  SET_TOKEN: (state, token) => {
    console.log('SET_TOKEN ' + token)
    state.token = token
    setToken(token)
  },
  SET_TOKENEXPIRE: (state, expire) => {
    state.tokenExpire = expire
    setTokenExpire(expire)
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  }
}

const actions = {
  // user login
  login ({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        if (response.code === 200) {
          commit('SET_TOKEN', response.token)
          commit('SET_TOKENEXPIRE', response.expire)
          resolve()
        } else {
          if (response.status === 401) {
            reject(response.data.message)
          }
        }
      }).catch(error => {
        console.log('login faild', error)
        alert('login faild' + error)
        reject(error)
      })
    })
  },
  signup ({ commit }, userInfo) {
    const { user_name, passwd, re_passwd } = userInfo
    return new Promise((resolve, reject) => {
      signUp({ user_name: user_name.trim(), passwd: passwd, re_passwd: re_passwd }).then(response => {
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  refreshToken ({ commit }) {
    return new Promise((resolve, reject) => {
      refreshToken().then(response => {
        const data = response
        commit('SET_TOKEN', data.token)
        commit('SET_TOKENEXPIRE', data.expire)
        setToken(data.token)
        setTokenExpire(data.expire)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  // get user info
  getInfo ({ commit, state }) {
    return new Promise((resolve, reject) => {
      commit('SET_TOKEN', getToken())
      commit('SET_TOKENEXPIRE', getTokenExpire())
      getInfo(getToken()).then(response => {
        console.log('getInfo response:' + JSON.stringify(response))
        const { data } = response
        if (!data) {
          alert('!data')
          reject('Verification failed, please Login again.')
        }
        const { roles, name, avatar, introduction } = data
        console.log(JSON.stringify(data))
        console.log(JSON.stringify(roles))
        // roles must be a non-empty array
        if (!roles || roles.length <= 0) {
          reject('getInfo: roles must be a non-null array!')
        }

        commit('SET_ROLES', roles)
        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        commit('SET_INTRODUCTION', introduction)
        resolve(data)
      }).catch(error => {
        alert('getInfo catch' + error)
        reject(error)
      })
    })
  },

  // user logout
  logout ({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      logout(getToken).then(() => {
        commit('SET_TOKEN', '')
        commit('SET_ROLES', [])
        removeToken()
        removeTokenExpire()
        resetRouter()

        // reset visited views and cached views
        // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
        dispatch('tagsView/delAllViews', null, { root: true })

        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },

  // remove token
  resetToken ({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_TOKENEXPIRE', '')
      removeToken()
      removeTokenExpire()
      resolve()
    })
  },

  // dynamically modify permissions
  async changeRoles ({ commit, dispatch }, role) {
    const token = role + '-token'

    commit('SET_TOKEN', token)
    setToken(token)

    const { roles } = await dispatch('getInfo')

    resetRouter()

    // generate accessible routes map based on roles
    const accessRoutes = await dispatch('permission/generateRoutes', roles, { root: true })
    // dynamically add accessible routes
    router.addRoutes(accessRoutes)

    // reset visited views and cached views
    dispatch('tagsView/delAllViews', null, { root: true })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
