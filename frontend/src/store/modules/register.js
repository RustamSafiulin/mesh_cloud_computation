
import axios from 'axios'

const api_url = `http://${window.location.hostname}:8081/api/v1/accounts`

export default {
    state: {
    },
    mutations: {
    },
    actions: {
        async register({commit}, signupInfo) {
            try {
                const response = await axios.post(api_url, signupInfo)
                const accountInfo = response.data
                return accountInfo
            } catch(e) {
                throw new Error(e);
            }
        }
    }
}