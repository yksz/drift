<app-open>
  <div>{ content }</div>
  <script>
    let path = location.pathname;
    path = path.substring('/open'.length);
    let self = this;
    let apiURL = '/api/open?path=' + path;
    fetch(apiURL)
    .then(response => response.text())
    .then(text => {
      self.content = text;
      self.update();
    });
  </script>
</app-open>
