
import axios from 'axios'

const api_url = `http://${window.location.hostname}:8081/api/v1/tasks`

export default {
    state: {
        tasks: []
    },
    mutations: {
        updateTasks(state, tasks) {
            state.tasks = tasks
        }
    },
    actions: {
        async createTask({commit}, task) {
            try {
                const response = await axios.post(api_url, task)
                const createdTask = response.data
                return createdTask
            } catch(e) {
                throw new Error(e);
            }
        },
        async deleteTask({commit}, id) {
            try {
                const response = axios.delete(api_url + `/${id}`)
            } catch(e) {
                throw new Error(e);
            }
        },
        async fetchTasks({commit}) {
            try {
                
                const response = await axios.get(api_url)
                const fectchedTasks = response.data
                
                commit("updateTasks", fectchedTasks);
            } catch (e) {
                throw new Error(e);
            }
        },
        async uploadTaskData({commit}, {taskId, formData, progressCallback}) {
            try {
                await axios.post(api_url + `/${taskId}` + '/upload', formData, {
                    onUploadProgress: progressCallback
                });
            } catch(e) {
                throw new Error(e);
            }
        }
    }, 
    getters: {
        allAccountTasks: state => state.tasks
    }
}