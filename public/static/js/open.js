const OPEN_PATH_PREFIX = '/open';

new Vue({
  el: '#open',

  data: {
    content: '',
  },

  created: function() {
    this.fetch();
  },

  methods: {
    fetch: function() {
      let path = location.pathname;
      path = path.substring(OPEN_PATH_PREFIX.length);
      let self = this;
      let apiURL = '/api/open?path=' + path;
      $.ajax(apiURL).done(data => {
        self.content = data;
      });
    }
  }
});
