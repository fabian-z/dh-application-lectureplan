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

    document.getElementById("submit").addEventListener("click", function() {
       // Example click handler
    }, false);

}, false);
