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

function now() {
    return moment().format('YYYY-MM-DDTHH:mm:ss\\Z')
}

export default {
    getDueString,
    now
}