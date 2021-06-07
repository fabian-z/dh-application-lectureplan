import "../css/fonts.css";

import 'bootstrap';
// Could use Bootstrap SCSS sources here
// see https://getbootstrap.com/docs/5.0/getting-started/webpack/#importing-styles
// precompiled CSS works for now
import 'bootstrap/dist/css/bootstrap.css';

import { Calendar } from '@fullcalendar/core';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import listPlugin from '@fullcalendar/list';
import interactionPlugin from '@fullcalendar/interaction';
import bootstrapPlugin from '@fullcalendar/bootstrap';
import scrollgridPlugin from '@fullcalendar/scrollgrid';
import "../css/main.css";


document.addEventListener('DOMContentLoaded', function() {
    // Setup handlers when DOM is ready

    // DOM event handlers

    let calendarEl = document.getElementById('calendar');
    let calendar = new Calendar(calendarEl, {
        plugins: [ dayGridPlugin, timeGridPlugin, listPlugin, interactionPlugin, bootstrapPlugin, scrollgridPlugin ],
        initialView: 'dayGridMonth',
        weekNumberCalculation: 'ISO',
        locales: [calendar_de],
        locale: 'de',
        themeSystem: 'bootstrap',
        headerToolbar: {
          left: 'prev,next today',
          center: 'title',
          right: 'dayGridMonth,timeGridWeek,listWeek',
        },
        bootstrapFontAwesome: {
            close: 'fa-times',
            prev: 'fa-chevron-left',
            next: 'fa-chevron-right',
            prevYear: 'fa-angle-double-left',
            nextYear: 'fa-angle-double-right',
          },
          
      });
    calendar.render();

    /*document.getElementById("submit").addEventListener("click", function() {
       // Example click handler
    }, false);*/

}, false);

// Translate here because import of locale submodule gives Webpack errors
// "Parsed request is a module" and "Field 'browser' doesn't contain a valid alias configuration"
var calendar_de = {
  code: 'de',
  week: {
    dow: 1, // Monday is the first day of the week.
    doy: 4, // The week that contains Jan 4th is the first week of the year.
  },
  buttonText: {
    prev: 'Zurück',
    next: 'Vor',
    today: 'Heute',
    year: 'Jahr',
    month: 'Monat',
    week: 'Woche',
    day: 'Tag',
    list: 'Terminübersicht',
  },
  weekText: 'KW',
  allDayText: 'Ganztägig',
  moreLinkText: function(n) {
    return '+ weitere ' + n;
  },
  noEventsText: 'Keine Ereignisse anzuzeigen',
};