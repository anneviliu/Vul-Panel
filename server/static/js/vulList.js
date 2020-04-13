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
        var color
        console.log(this.vulData);
        for (var i in this.vulData) {
            if (this.vulData[i].Read) {
                color = "#0090ff"
            } else {
                color = "rgba(255,0,104,0.63)"
            }
            this.recentList +=
                "                <div class=\"media text-muted pt-3\">" +
                "                    <p class=\"media-body pb-3 mb-0 small lh-125 border-bottom border-gray\">" +
                "                        <strong class=\"d-block text-gray-dark\">" + this.vulData[i].CreatedAt + "</strong>" +
                "<a style='color: " + color + ";font-weight:bold; text-decoration: none' href="+ this.vulData[i].Url + ">" + this.vulData[i].Host +"</a>" + "<p style='color: red;font-size: 12px;'>" +this.vulData[i].Title + "</p>" +
                "                </div>"
        }
    },
    methods:{
        async getRecentList() {
            return await axios.get('/api/recentList')
        },
    }

});