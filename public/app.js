new Vue({
    el: 'body',

    data: {
        tasks: [],
        newTask: {}
    },

    created() {
        this.$http.get('/tasks').then(response => {
            this.tasks = response.data.items || []
        })
    },

    methods: {
        createTask() {
            if (!this.newTask.name.trim()) {
                this.newTask = {}
                return
            }

            this.$http.put('/tasks', this.newTask).success(response => {
                this.newTask.id = response.created
                this.tasks.push(this.newTask)
                this.newTask = {}
            }).error(console.log);
        },

        deleteTask(index) {
            this.$http.delete(`/tasks/${this.tasks[index].id}`).success(response => {
                this.tasks.splice(index, 1)
            }).error(console.log)
        }
    }
})
