new Vue({
    el: '#app',
    delimiters: ["{[", "]}"], // 可自定义符号
    data() {
        return {
            test: "",
            VulList: "",
            vulData: this.vulData,
            totalItems: 0,
            totalPages: 0,
            perPage : 20,
            pageNow : 1,
            pageList : [],
        }
    },

    mounted() {
        this.loadPageList()
        if (localStorage.getItem("pageNow") == null) {
            this.pageNow = 0
        } else {
           this.pageNow = localStorage.getItem("pageNow")
        }
        this.loadVulList(this.pageNow)
    },
    methods:{
        async getListByPage(p) {
            return await axios.get('/api/getPages?p='+p)
        },
        async getTotalPages() {
            return await axios.get('/api/getPages?t=totalPages')
        },

        async loadVulList(page) {
            this.VulList = ""
            this.vulData = (await this.getListByPage(page)).data;
            var color = ""
            // console.log(this.vulData);
            for (var i in this.vulData) {
                if (this.vulData[i].Read) {
                    color = "#0090ff"
                } else {
                    color = "rgba(255,0,104,0.63)"
                }

                if (this.vulData[i].VulUrl.length > 70) {
                    var VulUrl = this.vulData[i].VulUrl.substring(0,70) + "...."
                } else {
                    VulUrl = this.vulData[i].VulUrl
                }
                this.VulList +=
                    "                <div class=\"media text-muted pt-3\">" +
                    "                    <p class=\"media-body pb-3 mb-0 small lh-125 border-bottom border-gray\">" +
                    "                        <strong class=\"d-block text-gray-dark\">" + this.vulData[i].CreatedAt + "</strong>" +
                    "<a style='color: " + color + ";font-weight:bold; text-decoration: none' href="+ this.vulData[i].Url + ">" + this.vulData[i].Host +"</a>" + "<a style='padding-left: 50px'>"+VulUrl + "</a>" + "<p style='color: red;font-size: 12px;'>" +this.vulData[i].Title + "</p>" +
                    "                </div>"
            }
        },
        
        switchToPage: function (pageNo) {
            if (pageNo < 0 || pageNo > this.totalPages+1){
                return false;
            }
            localStorage.setItem("pageNow",pageNo)
            this.pageNow = pageNo

            this.loadVulList(pageNo)
            this.loadPageList()
        },

        async loadPageList() {
            var limit = 20
            var i = 1
            this.pageList = [];
            this.totalPages = (await this.getTotalPages()).data;
            // if (this.pageNow > this.totalPages) {
            //     return false
            // }
            if (this.totalPages > limit && this.pageNow < limit-1) {
                for(i;i<=limit;i++) {
                    this.pageList.push(i)
                }
            }
            if (this.pageNow >= 19) {
                this.pageList = []
                console.log(this.pageList)
                for(i=this.pageNow-18;i<=this.pageNow +1 ;i++) {
                    this.pageList.push(i)
                }
            }
        }
    }

});

// function getListByPage(pageNow) {
//     $.ajax({
//         url:"/user/"+pageNow,
//         success:function (datas) {
//             vueApp.userList = datas.content;
//             vueApp.totalPages = datas.totalPages;
//             vueApp.pageNow = pageNow;
//         },
//         error:function (res) {
//             console.log(res);
//         }
//     });
// }