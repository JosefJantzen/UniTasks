import moment from 'moment';

/*function formatDate (value) {
    return moment(String(value)).format('DD.MM.YYYY hh:mm')
}*/

function getDueString(v) {
    let due = moment(String(v))
    if (due.diff(moment.now(), 'days') < 1) {
        return "Less than a day left"
    }
    return "You have " + due.diff(moment.now(), 'days') + " days"
}

export default {
    getDueString
}