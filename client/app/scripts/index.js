window.onload = function () {
    var id = window.location.pathname.split('/', 2)[1];
    var app = new Vue({
        el: '#app',
        data: {
            loading: true,
            listImgs: {},
            listMode: false,
            all: {shown: false},
            display_name: '',
            metadata: {},
            system: {},
            reports: {},
            showHideAll: function (all, reports) {
                all.shown = !all.shown;
                Object.keys(reports).forEach(r => {
                    reports[r].expanded_ = all.shown;
                });
            },
            goTo: function(imageId) {
                if (imageId === undefined) {
                    window.location.pathname = '/';
                } else {
                    window.location.pathname = '/' + imageId;
                }
            }
        }
    });

    if (id === undefined || id === '') {
        // 401
        // just kidding we want a list
        fetch('/reports/').then(response => {
            return response.json();
        }).then(json => {
            app.listMode = true;
            app.listImgs = json;
            app.loading = false;
        });
    } else {
        fetch('/reports/' + id).then(response => {
            return response.json();
        }).then(json => {
            var data = JSON.parse(json.report);
            Object.keys(data.reports).forEach(r => {
                data.reports[r].expanded_ = false;
            });

            app.display_name = data.system.hostname || data.system.system_id;
            app.metadata = data.system.metadata;
            app.system = data.system;
            app.reports = data.reports;
            app.listMode = false;
            app.loading = false;
        });
    }
};
