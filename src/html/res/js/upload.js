let submit = document.getElementById('submit')

let button = document.getElementById('upload')

submit.disabled = true
submit.classList += " disabled"

button.onchange = function () {
    document.getElementById('upload-filename').value = this.value;
    submit.disabled = false;
    submit.setAttribute('class', 'button button-default')
}
