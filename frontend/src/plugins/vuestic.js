import { createVuestic, createIconsConfig } from 'vuestic-ui'
import 'vuestic-ui/css'
import 'material-design-icons-iconfont/dist/material-design-icons.min.css' 

const vuestic = createVuestic({
    config: {
        colors: {
            variables: {
                primary: '#055052',
                secondary: '#53B8BB',
                success: '#40e583',
                info: '#f3f2c9',
                danger: '#e34b4a',
                warning: '#ffc200',
                dark: '#003638',
                light: '#def0f0',
                white: '#ffffff'
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