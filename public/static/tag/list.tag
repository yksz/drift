<app-list>
  <ul>
    <li each={ file in filelist }>
      <a href={ file.path }>{ file.name }</a>
    </li>
  </ul>
  <script>
    let path = location.pathname;
    path = path.substring('/list'.length);
    let self = this;
    let apiURL = '/api/list?path=' + path;
    fetch(apiURL)
    .then(response => response.json())
    .then(json => {
      json = json.filter(file => !file.is_hidden);
      json.forEach(file => {
        if (file.is_dir) {
          file.path = '/list/' + file.path;
        } else {
          file.path = '/open/' + file.path;
        }
      });
      self.filelist = json;
      self.update();
    });
  </script>
</app-list>
