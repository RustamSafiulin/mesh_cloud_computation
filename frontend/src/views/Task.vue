<template>
  <div>
    <div class="page-title">
      <h3>Задачи</h3>
    </div>

    <p class="center" v-if="!accountTasks.length">Нет задач.</p>

    <div class="table-wrapper" v-else>
      <table class="table table-striped table-hover">
        <thead>
          <tr>
            <th>Описание</th>
            <th>Состояние</th>
            <th>Данные</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(task, idx) in accountTasks" :key="idx">
            <td>{{task.description}}</td>
            <td>{{task.state}}</td>
            <td>
              <form enctype="multipart/form-data" @submit.prevent="onUploadTaskData(task.task_id)">
                <div class="fields">
                  <label>Upload file</label>
                  <input type="file" @change="onUploadFileSelect" />
                </div>
                <div class="fields">
                  <button>Submit</button>
                </div>
                <progress max="100" :value.prop="uploadPercentage"></progress>
              </form>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "task",
  data() {
    return {
      accountTasks: [],
      uploadFile: null,
      uploadPercentage: 0
    };
  },
  async mounted() {
    await this.fetchTasks();
  },
  methods: {
    notifyUploadProgress(progressEvent) {
      console.log("progress");
      this.uploadPercentage = parseInt(
        Math.round((progressEvent.loaded * 100) / progressEvent.total)
      );
    },
    onUploadFileSelect(e) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) return;

      this.uploadFile = files[0];
    },
    async onUploadTaskData(taskId) {
      const formData = new FormData();
      formData.append("task_data", this.uploadFile);

      var progressCallback = this.notifyUploadProgress;

      try {
        await this.$store.dispatch("uploadTaskData", {
          taskId,
          formData,
          progressCallback
        });

        alert("Uploaded!!!");
      } catch (e) {
        console.log(e);
      }
    },
    async fetchTasks() {
      try {
        await this.$store.dispatch("fetchTasks");
        console.log(this.$store.getters.allAccountTasks);
        this.accountTasks = this.$store.getters.allAccountTasks;
      } catch (e) {
        console.log(e);
      }
    }
  }
};
</script>

<style lang="scss" scoped>
table.table-striped tbody tr:nth-of-type(odd) {
  background-color: #fcfcfc;
}
table.table-striped.table-hover tbody tr:hover {
  background: #f5f5f5;
}

table.table tr th:first-child {
  width: 150px;
}

table {
  border-collapse: collapse;
  table-layout: auto;
  width: 100%;
}

thead th,
tbody td {
  text-align: center;
}

thead th {
  background-color: #f5f5f5;
}
</style>