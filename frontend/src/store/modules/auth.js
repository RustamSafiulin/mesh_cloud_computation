
import axios from 'axios'

const api_url = `http://${window.location.hostname}:8081/api/v1/accounts`
const user_info = JSON.parse(localStorage.getItem('user_info_3d_mesh'))

export default {
	state: {
        status: '',
		user: user_info
    },
	mutations: {
		auth_request(state){
	    	state.status = 'loading'
	  	},
	  	auth_success(state, user){

		    state.status = 'success'
		    state.user = user
	  	},
	  	auth_error(state){
	    	state.status = 'error'
	  	},
	  	logout(state){
	    	state.status = ''
	    	state.user = null
		}
	},
    actions: {
		async login({commit}, user) {
			try {
				commit('auth_request')
				
				const response = await axios.post(api_url + '/signin', user)
				const userInfo = response.data
				localStorage.setItem('user_info_3d_mesh', JSON.stringify(userInfo))
				
				axios.defaults.headers.common['Authorization'] = "Bearer_" + userInfo.session_token
				commit('auth_success', userInfo)
				
			} catch (e) {
				commit('auth_error')
				localStorage.removeItem('user_info_3d_mesh')
				throw new Error(e)
			}
		},
        logout({commit}){
		    return new Promise((resolve, reject) => {
		      	commit('logout')
				localStorage.removeItem('user_info_3d_mesh')
		      	delete axios.defaults.headers.common['Authorization']
		      	resolve()
		    })
	  	}
    },
	getters : {
		isLoggedIn(state) { 
			if (state.user) {
				return !!state.user.session_token
			}
			return false
		},
        authStatus(state) {
			return state.status
		},
		userName(state) {
			if (state.user) {
				return state.user.username
			}
			return ''
		},
		hasAdminRights(state) {
			if (state.user) {
				return state.user.is_admin
			}
			return false
		},
		currentAccountId(state) {
			if (state.user) {
				return state.user.id
			}
			return 0
		}
    }
}