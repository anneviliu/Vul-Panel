new Vue({
    el: '#app',
    delimiters: ["{[", "]}"], // 可自定义符号
    data() {
        return {
            // test: "",
            // vulList: "",
            vulData: [],
            totalItems: 0,
            totalPages: 0,
            perPage : 20,
            pageNow : 1,
            pageList : [],
            // vulUrlx :"",
            itemStyle:{
                color : "rgba(255,0,104,0.63)",
                fontWeight:"bold",
                textDecoration:"none",
            }
        }
    },

    mounted() {
        this.pageList = [];
        this.loadPageList()
        if (localStorage.getItem("pageNow") == null) {
            this.pageNow = 0
        } else {
           this.pageNow = localStorage.getItem("pageNow")
        }
        this.getVulList(this.pageNow)
    },
    methods:{
        async getListByPage(p) {
            return await axios.get('/api/getPages?p='+p)
        },

        async getTotalPages() {
            return await axios.get('/api/getPages?t=totalPages')
        },

        async getVulList(page) {
            this.vulData = (await this.getListByPage(page)).data;
        },

        async pinHigh(p) {
            return await axios.get('/api/getPages?p='+p)
        },

        deleteItem: function (id) {
            var timestamp = (new Date()).valueOf();
            let data = {"id":id,"timestamp":timestamp};
            axios.post('/api/deleteItems',data)
            location.reload(true)
        },

        pinStatus: function (id,status) {
            let data = {"id":id,"status":status}
            axios.post('/api/pinStatus',data)
            location.reload(true)
        },
        
        switchToPage: function (pageNo) {
            if (pageNo <= 0 || pageNo > this.totalPages+1){
                return false;
            }
            localStorage.setItem("pageNow",pageNo)
            this.pageNow = pageNo

            this.loadVulList(pageNo)
            this.loadPageList()
        },

        async loadPageList() {
            this.pageList = [];
            var limit = 20
            var i = 1

            this.totalPages = (await this.getTotalPages()).data;
            if (Number(this.pageNow) > this.totalPages) {
                return false
            }

            if (this.totalPages > limit && Number(this.pageNow) < limit-1) {
                for(i;i<=limit;i++) {
                    this.pageList.push(i)
                }
            } else {
                for (i;i<=this.totalPages;i++) {
                    this.pageList.push(i)
                }
            }

            if (Number(this.pageNow) >= 19) {
                this.pageList = []
                for(var j=Number(this.pageNow)-18;j<=Number(this.pageNow)+1 ;j++) {
                    if (Number(this.pageNow) +1 > this.totalPages) {
                        return false
                    }
                    this.pageList.push(j)
                }
            }
        },
    }
});