new Vue({
    el: '#app',
    data() {
        return {
            email: "",
            password: "",
            inviteCode:"",
            username:"",
        }
    },
    methods: {
        register() {
            let data = {"email":this.email,"password":this.password,"username":this.username,"inviteCode":this.inviteCode};
            axios.post('/api/reg', data)
                .then(function (response) {
                    console.log(response);
                    if (response.data.msg == "注册成功"){
                        alert(response.data.msg);
                        location.href= "/login"
                    }
                    else {
                        alert(response.data.msg);
                    }
                })
                .catch(function (error) {
                    console.log(error);
                });
        }
    }
});