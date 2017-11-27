window.onload = function () {
    var id = window.location.pathname.split('/', 2)[1];
    console.log(id);
    if (id === undefined || id === '') {
        // 401
    } else {
        fetch('/report/' + id).then(response => {
            return response.json();
        }).then(json => {
            var data = JSON.parse(json.report);
            Object.keys(data.reports).forEach(r => {
                data.reports[r].expanded_ = false;
            });

            var app = new Vue({
                el: '#app',
                data: {
                    all: {shown: false},
                    display_name: data.system.hostname || data.system.system_id,
                    metadata: data.system.metadata,
                    system: data.system,
                    reports: data.reports,
                    showHideAll: function (all, reports) {
                        all.shown = !all.shown;
                        Object.keys(reports).forEach(r => {
                            reports[r].expanded_ = all.shown;
                        });
                    }
                }
            });
        });
    }
};
