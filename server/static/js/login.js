new Vue({
    el: '#app',
    data() {
        return {
            email: "",
            password: ""
        }
    },
    methods: {
        login() {
            let data = {"email":this.email,"password":this.password};
            axios.post('/api/login', data)
                .then(function (response) {
                    console.log(response);
                    if (response.data.msg == "登录成功"){
                        alert(response.data.msg);
                        location.href= "/"
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