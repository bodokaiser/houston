import {createReducer} from 'redux-create-reducer'

export default createReducer({
  cpu: { name: 'CPU Usage', value: '50%' },
  ram: { name: 'RAM Usage', value: '60%' },
  disk: { name: 'Disk Usage', value: '30%' },
  uptime: { name: 'Uptime', value: '30h' }
}, {})
