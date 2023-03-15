import moment from 'moment';

function formatDate (value) {
    if (value) {
        return moment(String(value)).format('DD.MM.YYYY hh:mm')
    }
}

export default {
    formatDate
}