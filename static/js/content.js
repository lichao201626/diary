function save() {
  var content = document.getElementById('text').value;
  axios({
    method: 'post',
    url: 'http://127.0.0.1:8880/save',
    data: {
      username: 'Fred',
      password: 'Flintstone',
      content
    }
  }).then(function(response) {
    window.alert("success")
  });
}

function rand() {
  var username = 'Fred';
  var password = 'Flintstone'
  window.location.href=`http://127.0.0.1:8880/rand?username=${username}&password=${password}`;
/*   axios({
    method: 'get',
    url: `http://127.0.0.1:8880/rand?username=${username}&password=${password}`,
    responseType: 'string'
  }).then(function (response) {
    // response.data.pipe(fs.createWriteStream('ada_lovelace.jpg'))
    // window.alert(response.data)
    // document.getElementById('text').value = response.data;
  }); */
}

function create() {
  var username = 'Fred';
  var password = 'Flintstone'
  window.location.href=`http://127.0.0.1:8880/write?username=${username}&password=${password}`;
/*   axios({
    method: 'get',
    url: `http://127.0.0.1:8880/write?username=${username}&password=${password}`,
    responseType: 'string'
  }) */
}
