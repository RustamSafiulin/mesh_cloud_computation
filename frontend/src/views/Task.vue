<template>
  <div>
    <div class="page-title">
      <h3>Задачи</h3>
    </div>

    <section>
      <form class="row" @submit.prevent="addNewTask">
        <div class="input-field col s4" id="addTaskInput">
          <input id="task_description" type="text" v-model="task_description" />
          <label for="task_description">Описание</label>
          <span class="helper-text">Введите описание задачи</span>
        </div>

        <button
          id="addTaskBtn"
          title="Добавить задачу"
          data-toggle="tooltip"
          class="btn waves-effect waves-light col s1"
          :disabled="$v.$invalid"
          type="submit"
        >
          <i class="material-icons right">add</i>
        </button>
      </form>
      <div
        class="error"
        v-if="$v.task_description.$touch()"
        id="task_error_label"
      >Описание Задачи не может быть пустым</div>

      <div>
        <div class="modal" ref="taskDeleteDialog" id="task_delete_confirm_modal">
          <div class="modal-content">
            <h4>Удаление информации о задаче</h4>
            <p>Вы действительно хотите удалить задачу '{{taskDeleteDialogData.text}}'?</p>
          </div>
          <div class="modal-footer">
            <button
              class="btn waves-effect waves-light btn-small"
              @click="confirmDeleteTask"
              id="task_delete_confirm_btn"
            >
              <i class="material-icons"></i>ОК
            </button>
            <button class="btn waves-effect waves-light btn-small" @click="cancelConfirmDeleteTask">
              <i class="material-icons right">cancel</i>Отмена
            </button>
          </div>
        </div>

        <p class="center" v-if="!accountTasks.length">Нет задач.</p>

        <div class="table-wrapper" v-else>
          <table class="table table-striped table-hover">
            <thead>
              <tr>
                <th>Описание</th>
                <th>Состояние</th>
                <th>Данные</th>
                <th>Действия</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="(info, idx) in accountTasks" :key="idx">
                <td>{{info.task.description}}</td>
                <td>{{taskStateDescription(info.task.state)}}</td>
                <td id="uploadForm">
                    <span class="btn btn-file">
                      <i class="material-icons left">cloud_upload</i>
                      Открыть
                      <input type="file" @change="onUploadFileSelect($event, info)"/>
                    </span>
                    <div class="row" id="uploadProgressBar">
                      <div class="progress col md2">
                        <div class="determinate" :style="info.percentageStyle"></div>
                      </div>
                      <div class="col s3" id="percentageLabel">
                        <label>{{info.uploadPercentage}}%</label>
                      </div>
                    </div>
                </td>
                <td>
                  <a
                    href="#"
                    class="delete"
                    title="Удалить"
                    data-toggle="tooltip"
                    v-on:click="showDeleteTaskConfirmModal(info.task, idx)"
                  >
                    <i class="material-icons">remove_circle_outline</i>
                  </a>
                  <a v-if="info.task.state === 0" href="#" class="start_task" title="Запустить" data-toggle="tooltip">
                    <i class="material-icons">play_arrow</i>
                  </a>
                  <a v-if="info.task.state === 1" href="#" class="stop_task" title="Запустить" data-toggle="tooltip">
                    <i class="material-icons">stop</i>
                  </a>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { required } from "vuelidate/lib/validators";

export default {
  name: "task",
  data() {
    return {
      task_description: "",
      accountTasks: [],

      taskDeleteDialog: null,
      taskDeleteDialogData: {
        text: "",
        task: null,
        viewIndex: 0
      }
    };
  },
  validations: {
    task_description: {
      required
    }
  },
  computed: {
    taskStateDescription() {
      return state => {
        if (state === 0) {
          return "Создана";
        } else if (state === 1) {
          return "Запущена";
        } else if (state === 2) {
          return "Выполнена";
        } else if (state === 3) {
          return "Остановлена";
        }

        return state;
      };
    }
  },
  async mounted() {
    await this.fetchTasks();

    this.taskDeleteDialog = M.Modal.init(this.$refs.taskDeleteDialog, {
      startingTop: "50%",
      endingTop: "40%"
    });
  },
  beforeDestroy() {
    if (this.taskDeleteDialog && this.taskDeleteDialog.destroy) {
      this.taskDeleteDialog.destroy();
    }
  },
  methods: {
    showDeleteTaskConfirmModal(task, idx) {
      this.taskDeleteDialogData.text = task.description;
      this.taskDeleteDialogData.task = task;
      this.taskDeleteDialogData.viewIndex = idx;
      this.taskDeleteDialog.open();
    },
    confirmDeleteTask() {
      this.deleteTask(
        this.taskDeleteDialogData.task.id,
        this.taskDeleteDialogData.viewIndex
      );
      this.taskDeleteDialog.close();
    },
    cancelConfirmDeleteTask() {
      this.taskDeleteDialog.close();
    },
    async onUploadFileSelect(e, info) {
      var files = e.target.files || e.dataTransfer.files;
      if (!files.length) return;

      info.uploadFile = files[0];
      await this.onUploadTaskData(info);
    },
    uploadPercentageStyle(valuePercentage) {
      return {
        width: `${valuePercentage}%`
      };
    },
    async onUploadTaskData(info) {
      const formData = new FormData();
      formData.append("task_data", info.uploadFile);

      var taskId = info.task.id;
      var progressCallback = function(progressEvent) {
        info.uploadPercentage = parseInt(
          Math.round((progressEvent.loaded * 100) / progressEvent.total)
        );
      };

      try {
        await this.$store.dispatch("uploadTaskData", {
          taskId,
          formData,
          progressCallback
        });
      } catch (e) {
        console.log(e);
      }
    },
    async fetchTasks() {
      try {
        await this.$store.dispatch("fetchTasks");

        const tasks = this.$store.getters.allAccountTasks;
        let accountTasks = [];

        tasks.forEach(function(t) {
          let taskInfo = {};
          taskInfo.task = t;
          taskInfo.uploadFile = null;
          taskInfo.uploadPercentage = 0;
          accountTasks.push(taskInfo);
        });

        this.accountTasks = accountTasks;
      } catch (e) {
        console.log(e);
      }
    },
    async addNewTask() {
      try {
        const description = this.task_description;
        const createdTask = await this.$store.dispatch("createTask", {
          description
        });

        var taskInfo = {
          task: createdTask,
          uploadFile: null,
          uploadPercentage: 0,
          percentageStyle: uploadPercentageStyle(0)
        };

        this.accountTasks.push(taskInfo);
      } catch (e) {
        console.log(e);
      }
    },
    async deleteTask(id, index) {
      try {
        await this.$store.dispatch("deleteTask", id);
        this.accountTasks.splice(index, 1);
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

table.table td a {
  font-weight: bold;
  color: #566787;
  display: inline-block;
  text-decoration: none;
}
table.table td a:hover {
  color: #2196f3;
}

table.table td a.delete {
  color: #f44336;
}

table.table td a.stop_task {
  color: #f44336;
}

table.table td a.start_task {
  color: #26a699;
}

thead th,
tbody td {
  text-align: center;
}

thead th {
  background-color: #f5f5f5;
}

#task_delete_confirm_btn {
  margin-right: 10px;
}

#addTaskBtn {
  max-width: 40px;
  margin-top: 20px;
}

#addTaskInput {
  max-width: 300px;
}

#task_delete_confirm_modal {
  max-width: 600px;
  max-height: 200px;
}

#task_error_label {
  margin-top: -20px;
}

#uploadForm {
  width: 180px;
}

#uploadProgressBar {
  width: 150px;
}

#percentageLabel {
  margin-top: -10px;
}

.btn-file {
  position: relative;
  overflow: hidden;
}
.btn-file input[type="file"] {
  position: absolute;
  top: 0;
  right: 0;
  min-width: 100%;
  min-height: 100%;
  font-size: 100px;
  text-align: right;
  filter: alpha(opacity=0);
  opacity: 0;
  outline: none;
  background: white;
  cursor: inherit;
  display: block;
}
</style>