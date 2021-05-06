// ES6 module imports here
import 'bootstrap';

// Webpack CSS imports
// Could use Bootstrap SCSS sources here
// see https://getbootstrap.com/docs/5.0/getting-started/webpack/#importing-styles
// precompiled CSS works for now
import 'bootstrap/dist/css/bootstrap.css';

import "../css/main.css";

document.addEventListener('DOMContentLoaded', function() {
    // Setup handlers when DOM is ready

    // DOM event handlers

    document.getElementById("submit").addEventListener("click", function() {
       // Example click handler
    }, false);

}, false);
