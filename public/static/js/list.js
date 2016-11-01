new Vue({
  el: '#list',

  data: {
    filelist: {}
  },

  created: function() {
    this.fetch();
  },

  methods: {
    fetch: function() {
      let path = location.pathname;
      if (path.startsWith('/list')) {
        path = path.substring('/list'.length);
      }
      let self = this;
      const apiURL = '/api/file?path=' + path;
      $.getJSON(apiURL).done(function(json) {
        json.forEach(function(file) {
          file.path = 'list/' + file.path;
        });
        self.filelist = json;
      });
    }
  }
})
