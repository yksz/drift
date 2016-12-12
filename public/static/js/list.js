const LIST_PATH_PREFIX = '/list';
const OPEN_PATH_PREFIX = '/open';

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
      path = path.substring(LIST_PATH_PREFIX.length);
      let self = this;
      let apiURL = '/api/list?path=' + path;
      $.getJSON(apiURL).done(json => {
        json = json.filter(file => !file.is_hidden);
        json.forEach(file => {
          if (file.is_dir) {
            file.path = LIST_PATH_PREFIX + '/' + file.path;
          } else {
            file.path = OPEN_PATH_PREFIX + '/' + file.path;
          }
        });
        self.filelist = json;
      });
    }
  }
});
