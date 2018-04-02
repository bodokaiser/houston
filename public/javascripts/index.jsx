import {h, render} from 'preact'

import App from './components/app'

let content = document.querySelector('content')

render(<App /> , content, content.firstElementChild)

if (module.hot) module.hot.accept()
