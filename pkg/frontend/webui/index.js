function strip(s) {
  return s.replace(/\s+/g, '');
}

async function echoResponse(endpoint, query, aborter) {
  const response = await fetch(`/${endpoint}?q=${query}`, {signal: aborter});
  const text = await response.text();
  if (response.ok) {
    return text;
  } else {
    throw new Error(text);
  }
}

async function add(endpoint, query, aborter) {
  const response = await fetch(`/${endpoint}?q=${query}`, {signal: aborter});
  const text = await response.text();
  if (response.ok) {
    return text;
  } else {
    throw new Error(text);
  }
}

function main() {
  const query = document.getElementById('pass');
  const addButton = document.getElementById('add');
  const echoButton = document.getElementById('echo');
  const loader = document.getElementById('test');
  const timer = document.getElementById('time');
  const status = document.getElementById('status');

  const inputs = document.querySelectorAll('input');

  inputs.forEach(input => {
    input.setAttribute('autocomplete', 'off')
    input.setAttribute('autocorrect', 'off')
    input.setAttribute('autocapitalize', 'off')
    input.setAttribute('spellcheck', false)
  })

  addButton.disabled = true;
  echoButton.disabled = true;
  status.innerText = location.hostname;

  let controller; 
  let pending = 0; 

  const add_user = () => {
    if(controller != undefined) {
      controller.abort();
    }
    controller = new AbortController();

    for (const endpoint of ['add']) {
      if (pending == 0) {
        loader.hidden = false;
      }
      pending++;
      var start_time = performance.now();
      add(endpoint, query.value, controller.signal).then((x) => {
        const results = JSON.parse(strip(x));
        if (results == null || results.length == 0) {
          loader.innerText = "...";
          timer.innerText = "...";
        } else {
          loader.innerText = x.replace(/"/g, "");
        }
      }).finally(() => {
        pending--;
        if (pending == 0) {
        }
      });
      timer.innerText = (performance.now() - start_time) + "ms";
    }
  }

  const echo = () => {
    if (controller != undefined) {
      controller.abort();
    }
    controller = new AbortController();

    for (const endpoint of ['echo']) {
      if (pending == 0) {
        loader.hidden = false;
      }
      pending++;

      var start_time = performance.now();
      echoResponse(endpoint, query.value, controller.signal).then((v) => {
        const results = JSON.parse(v);
        if (results == null || results.length == 0) {
          loader.innerText = "...";
          timer.innerText = "...";
        } else {
          loader.innerText = "vivian: " + v.replace(/"/g, "");
        }
      }).finally(() => {
        pending--;
        if (pending == 0) {
        }
      });
      timer.innerText = (performance.now() - start_time) + "ms";
    }
  }

  addButton.addEventListener('click', add_user);
  echoButton.addEventListener('click', echo);

  query.addEventListener('keypress', (e) => {
    if (e.key == 'Enter' && strip(query.value) != "") {
      echo();
    }
  });

  query.addEventListener('input', (e) => {
    if (strip(query.value) == "") {
      addButton.disabled = true;
      echoButton.disabled = true;
			loader.innerText = "...";
			timer.innerText = "...";
    } else {
      addButton.disabled = false;
      echoButton.disabled = false;
    }
  });
}

document.addEventListener('DOMContentLoaded', main);
