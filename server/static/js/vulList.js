new Vue({
    el: '#app',
    data() {
        return {
            test: "",
            recentList: "",
            vulData: this.vulData
        }
    },
    async created() {
        this.vulData = (await this.getRecentList()).data;
        console.log(this.vulData);
        for (var i in this.vulData) {
            this.recentList +=
                "                <div class=\"media text-muted pt-3\">" +
                "                    <p class=\"media-body pb-3 mb-0 small lh-125 border-bottom border-gray\">" +
                "                        <strong class=\"d-block text-gray-dark\">" + this.vulData[i].CreatedAt + "</strong>" +
                "<a style='color: rgba(255,0,104,0.63) ;font-weight:bold; text-decoration: none' href="+ this.vulData[i].Url + ">" + this.vulData[i].Host +"</a>" + "<p style='color: red'>" +this.vulData[i].Title + "</p>" +
                "                </div>"
        }
    },
    methods:{
        async getRecentList() {
            return await axios.get('/api/recentList')
        },
    }

});