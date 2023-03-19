import moment from 'moment';

/*function formatDate (value) {
    return moment(String(value)).format('DD.MM.YYYY hh:mm')
}*/

function getDueString(v) {
    let due = moment(String(v))
    if (due.diff(moment.utc(), 'days') < 0) {
        return "You are late"
    }
    else if (due.diff(moment.utc(), 'days') < 1) {
        return "Less than a day left"
    }
    return "You have " + due.diff(moment.now(), 'days') + " days"
}

function formatTimestamp(v) {
    return moment.utc(v).format('DD.MM.YYYY [at] HH:mm')
}

function now() {
    return moment.utc().format('YYYY-MM-DDTHH:mm:ss\\Z')
}

export default {
    getDueString,
    now,
    formatTimestamp
}