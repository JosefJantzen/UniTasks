import { createVuestic, createIconsConfig } from 'vuestic-ui'
import 'vuestic-ui/css'
import 'material-design-icons-iconfont/dist/material-design-icons.min.css' 

const vuestic = createVuestic({
    config: {
        colors: {
            variables: {
                primary: '#F3F2C9',
                secondary: '#53B8BB',
                success: '#40e583',
                info: '#2c82e0',
                danger: '#e34b4a',
                warning: '#ffc200',
                gray: '#055052',
                dark: '#003638',
                background: '#def0f0'
            }
        },
        icons: createIconsConfig({
            aliases: [
              {
                "name": "bell",
                "color": "#FFD43A",
                "to": "mdi-bell"
              },
              {
                "name": "ru",
                "to": "flag-icon-ru small"
              },
            ],
            fonts: [
              {
                name: 'mdi-{iconName}',
                resolve: ({iconName}) => ({ 
                    class: `material-icons`, 
                    content: iconName,
                    tag: 'span'
                }),
              },
              {
                name: 'flag-icon-{countryCode} {flagSize}',
                resolve: ({countryCode, flagSize}) => ({
                  class: `flag-icon flag-icon-${countryCode} flag-icon-${flagSize}`
                }),
              }
            ],
          })
    }
})

export default vuestic