<template>
    <div class="row">
      <div class="col-md-8">
        <div class="wrapper">
          <div class="col-md-2">
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="assignNumber">
              <span> Assing Number</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="add">
              <span> Add</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="sub">
              <span> Sub</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="mult">
              <span> Mult</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="div">
              <span> Div</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="ifelse">
              <span> If/Else</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="print">
              <span> Print</span>
            </div>
            
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="forIn">
              <span> For In</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="code">
              <span>Code</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="aws">
              <svg class="svg-inline--fa fa-aws fa-w-20" aria-hidden="true" focusable="false" data-prefix="fab" data-icon="aws" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512" data-fa-i2svg=""><path fill="currentColor" d="M180.41 203.01c-.72 22.65 10.6 32.68 10.88 39.05a8.164 8.164 0 0 1-4.1 6.27l-12.8 8.96a10.66 10.66 0 0 1-5.63 1.92c-.43-.02-8.19 1.83-20.48-25.61a78.608 78.608 0 0 1-62.61 29.45c-16.28.89-60.4-9.24-58.13-56.21-1.59-38.28 34.06-62.06 70.93-60.05 7.1.02 21.6.37 46.99 6.27v-15.62c2.69-26.46-14.7-46.99-44.81-43.91-2.4.01-19.4-.5-45.84 10.11-7.36 3.38-8.3 2.82-10.75 2.82-7.41 0-4.36-21.48-2.94-24.2 5.21-6.4 35.86-18.35 65.94-18.18a76.857 76.857 0 0 1 55.69 17.28 70.285 70.285 0 0 1 17.67 52.36l-.01 69.29zM93.99 235.4c32.43-.47 46.16-19.97 49.29-30.47 2.46-10.05 2.05-16.41 2.05-27.4-9.67-2.32-23.59-4.85-39.56-4.87-15.15-1.14-42.82 5.63-41.74 32.26-1.24 16.79 11.12 31.4 29.96 30.48zm170.92 23.05c-7.86.72-11.52-4.86-12.68-10.37l-49.8-164.65c-.97-2.78-1.61-5.65-1.92-8.58a4.61 4.61 0 0 1 3.86-5.25c.24-.04-2.13 0 22.25 0 8.78-.88 11.64 6.03 12.55 10.37l35.72 140.83 33.16-140.83c.53-3.22 2.94-11.07 12.8-10.24h17.16c2.17-.18 11.11-.5 12.68 10.37l33.42 142.63L420.98 80.1c.48-2.18 2.72-11.37 12.68-10.37h19.72c.85-.13 6.15-.81 5.25 8.58-.43 1.85 3.41-10.66-52.75 169.9-1.15 5.51-4.82 11.09-12.68 10.37h-18.69c-10.94 1.15-12.51-9.66-12.68-10.75L328.67 110.7l-32.78 136.99c-.16 1.09-1.73 11.9-12.68 10.75h-18.3zm273.48 5.63c-5.88.01-33.92-.3-57.36-12.29a12.802 12.802 0 0 1-7.81-11.91v-10.75c0-8.45 6.2-6.9 8.83-5.89 10.04 4.06 16.48 7.14 28.81 9.6 36.65 7.53 52.77-2.3 56.72-4.48 13.15-7.81 14.19-25.68 5.25-34.95-10.48-8.79-15.48-9.12-53.13-21-4.64-1.29-43.7-13.61-43.79-52.36-.61-28.24 25.05-56.18 69.52-55.95 12.67-.01 46.43 4.13 55.57 15.62 1.35 2.09 2.02 4.55 1.92 7.04v10.11c0 4.44-1.62 6.66-4.87 6.66-7.71-.86-21.39-11.17-49.16-10.75-6.89-.36-39.89.91-38.41 24.97-.43 18.96 26.61 26.07 29.7 26.89 36.46 10.97 48.65 12.79 63.12 29.58 17.14 22.25 7.9 48.3 4.35 55.44-19.08 37.49-68.42 34.44-69.26 34.42zm40.2 104.86c-70.03 51.72-171.69 79.25-258.49 79.25A469.127 469.127 0 0 1 2.83 327.46c-6.53-5.89-.77-13.96 7.17-9.47a637.37 637.37 0 0 0 316.88 84.12 630.22 630.22 0 0 0 241.59-49.55c11.78-5 21.77 7.8 10.12 16.38zm29.19-33.29c-8.96-11.52-59.28-5.38-81.81-2.69-6.79.77-7.94-5.12-1.79-9.47 40.07-28.17 105.88-20.1 113.44-10.63 7.55 9.47-2.05 75.41-39.56 106.91-5.76 4.87-11.27 2.3-8.71-4.1 8.44-21.25 27.39-68.49 18.43-80.02z"></path></svg><!-- <i class="fab fa-aws"></i> --><span> AWS</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="log">
              <svg class="svg-inline--fa fa-file-signature fa-w-18" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="file-signature" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512" data-fa-i2svg=""><path fill="currentColor" d="M218.17 424.14c-2.95-5.92-8.09-6.52-10.17-6.52s-7.22.59-10.02 6.19l-7.67 15.34c-6.37 12.78-25.03 11.37-29.48-2.09L144 386.59l-10.61 31.88c-5.89 17.66-22.38 29.53-41 29.53H80c-8.84 0-16-7.16-16-16s7.16-16 16-16h12.39c4.83 0 9.11-3.08 10.64-7.66l18.19-54.64c3.3-9.81 12.44-16.41 22.78-16.41s19.48 6.59 22.77 16.41l13.88 41.64c19.75-16.19 54.06-9.7 66 14.16 1.89 3.78 5.49 5.95 9.36 6.26v-82.12l128-127.09V160H248c-13.2 0-24-10.8-24-24V0H24C10.7 0 0 10.7 0 24v464c0 13.3 10.7 24 24 24h336c13.3 0 24-10.7 24-24v-40l-128-.11c-16.12-.31-30.58-9.28-37.83-23.75zM384 121.9c0-6.3-2.5-12.4-7-16.9L279.1 7c-4.5-4.5-10.6-7-17-7H256v128h128v-6.1zm-96 225.06V416h68.99l161.68-162.78-67.88-67.88L288 346.96zm280.54-179.63l-31.87-31.87c-9.94-9.94-26.07-9.94-36.01 0l-27.25 27.25 67.88 67.88 27.25-27.25c9.95-9.94 9.95-26.07 0-36.01z"></path></svg><!-- <i class="fas fa-file-signature"></i> --><span> File Log</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="google">
              <svg class="svg-inline--fa fa-google-drive fa-w-16" aria-hidden="true" focusable="false" data-prefix="fab" data-icon="google-drive" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M339 314.9L175.4 32h161.2l163.6 282.9H339zm-137.5 23.6L120.9 480h310.5L512 338.5H201.5zM154.1 67.4L0 338.5 80.6 480 237 208.8 154.1 67.4z"></path></svg><!-- <i class="fab fa-google-drive"></i> --><span> Google Drive save</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="email">
              <svg class="svg-inline--fa fa-at fa-w-16" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="at" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M256 8C118.941 8 8 118.919 8 256c0 137.059 110.919 248 248 248 48.154 0 95.342-14.14 135.408-40.223 12.005-7.815 14.625-24.288 5.552-35.372l-10.177-12.433c-7.671-9.371-21.179-11.667-31.373-5.129C325.92 429.757 291.314 440 256 440c-101.458 0-184-82.542-184-184S154.542 72 256 72c100.139 0 184 57.619 184 160 0 38.786-21.093 79.742-58.17 83.693-17.349-.454-16.91-12.857-13.476-30.024l23.433-121.11C394.653 149.75 383.308 136 368.225 136h-44.981a13.518 13.518 0 0 0-13.432 11.993l-.01.092c-14.697-17.901-40.448-21.775-59.971-21.775-74.58 0-137.831 62.234-137.831 151.46 0 65.303 36.785 105.87 96 105.87 26.984 0 57.369-15.637 74.991-38.333 9.522 34.104 40.613 34.103 70.71 34.103C462.609 379.41 504 307.798 504 232 504 95.653 394.023 8 256 8zm-21.68 304.43c-22.249 0-36.07-15.623-36.07-40.771 0-44.993 30.779-72.729 58.63-72.729 22.292 0 35.601 15.241 35.601 40.77 0 45.061-33.875 72.73-58.161 72.73z"></path></svg><!-- <i class="fas fa-at"></i> --><span> Email send</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="template">
              <svg class="svg-inline--fa fa-code fa-w-20" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="code" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 640 512" data-fa-i2svg=""><path fill="currentColor" d="M278.9 511.5l-61-17.7c-6.4-1.8-10-8.5-8.2-14.9L346.2 8.7c1.8-6.4 8.5-10 14.9-8.2l61 17.7c6.4 1.8 10 8.5 8.2 14.9L293.8 503.3c-1.9 6.4-8.5 10.1-14.9 8.2zm-114-112.2l43.5-46.4c4.6-4.9 4.3-12.7-.8-17.2L117 256l90.6-79.7c5.1-4.5 5.5-12.3.8-17.2l-43.5-46.4c-4.5-4.8-12.1-5.1-17-.5L3.8 247.2c-5.1 4.7-5.1 12.8 0 17.5l144.1 135.1c4.9 4.6 12.5 4.4 17-.5zm327.2.6l144.1-135.1c5.1-4.7 5.1-12.8 0-17.5L492.1 112.1c-4.8-4.5-12.4-4.3-17 .5L431.6 159c-4.6 4.9-4.3 12.7.8 17.2L523 256l-90.6 79.7c-5.1 4.5-5.5 12.3-.8 17.2l43.5 46.4c4.5 4.9 12.1 5.1 17 .6z"></path></svg><!-- <i class="fas fa-code"></i> --><span> Template</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="multiple">
              <svg class="svg-inline--fa fa-code-branch fa-w-12" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="code-branch" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" data-fa-i2svg=""><path fill="currentColor" d="M384 144c0-44.2-35.8-80-80-80s-80 35.8-80 80c0 36.4 24.3 67.1 57.5 76.8-.6 16.1-4.2 28.5-11 36.9-15.4 19.2-49.3 22.4-85.2 25.7-28.2 2.6-57.4 5.4-81.3 16.9v-144c32.5-10.2 56-40.5 56-76.3 0-44.2-35.8-80-80-80S0 35.8 0 80c0 35.8 23.5 66.1 56 76.3v199.3C23.5 365.9 0 396.2 0 432c0 44.2 35.8 80 80 80s80-35.8 80-80c0-34-21.2-63.1-51.2-74.6 3.1-5.2 7.8-9.8 14.9-13.4 16.2-8.2 40.4-10.4 66.1-12.8 42.2-3.9 90-8.4 118.2-43.4 14-17.4 21.1-39.8 21.6-67.9 31.6-10.8 54.4-40.7 54.4-75.9zM80 64c8.8 0 16 7.2 16 16s-7.2 16-16 16-16-7.2-16-16 7.2-16 16-16zm0 384c-8.8 0-16-7.2-16-16s7.2-16 16-16 16 7.2 16 16-7.2 16-16 16zm224-320c8.8 0 16 7.2 16 16s-7.2 16-16 16-16-7.2-16-16 7.2-16 16-16z"></path></svg><!-- <i class="fas fa-code-branch"></i> --><span> Multiple inputs/outputs</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="personalized">
              <svg class="svg-inline--fa fa-fill fa-w-16" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="fill" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M502.63 217.06L294.94 9.37C288.69 3.12 280.5 0 272.31 0s-16.38 3.12-22.62 9.37l-81.58 81.58L81.93 4.77c-6.24-6.25-16.38-6.25-22.62 0L36.69 27.38c-6.24 6.25-6.24 16.38 0 22.63l86.19 86.18-94.76 94.76c-37.49 37.49-37.49 98.26 0 135.75l117.19 117.19c18.75 18.74 43.31 28.12 67.87 28.12 24.57 0 49.13-9.37 67.88-28.12l221.57-221.57c12.49-12.5 12.49-32.76 0-45.26zm-116.22 70.97H65.93c1.36-3.84 3.57-7.98 7.43-11.83l13.15-13.15 81.61-81.61 58.61 58.6c12.49 12.49 32.75 12.49 45.24 0 12.49-12.49 12.49-32.75 0-45.24l-58.61-58.6 58.95-58.95 162.45 162.44-48.35 48.34z"></path></svg><!-- <i class="fas fa-fill"></i> --><span> Personalized</span>
            </div>
            <div class="drag-drawflow" draggable="true" v-on:dragstart="drag" data-node="dbclick">
              <svg class="svg-inline--fa fa-mouse fa-w-12" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="mouse" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 384 512" data-fa-i2svg=""><path fill="currentColor" d="M0 352a160 160 0 0 0 160 160h64a160 160 0 0 0 160-160V224H0zM176 0h-16A160 160 0 0 0 0 160v32h176zm48 0h-16v192h176v-32A160 160 0 0 0 224 0z"></path></svg><!-- <i class="fas fa-mouse"></i> --><span> DBClick!</span>
            </div>
          </div>
          <div class="col-md-10">
            <div class="menu">
              <ul>
                <li>
                  <button type="button" v-on:click="clear" class="btn btn-outline-danger">Clear</button>
                </li>
                <li>
                  <button type="button" v-on:click="parse" class="btn btn-outline-secondary">Parse</button>
                </li>
                <li>
                  <button type="button" v-on:click="save" class="btn btn-outline-primary">Save</button>
                </li>
              </ul>
            </div>
            <div id="drawflow" v-on:drop="drop" v-on:dragover="allowDrop" class="parent-drawflow" tabindex="0">
              <div class="btn-lock">
                <svg id="lock" class="svg-inline--fa fa-lock fa-w-14" onclick="editor.editor_mode='fixed'; changeMode('lock');" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="lock" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 448 512" data-fa-i2svg=""><path fill="currentColor" d="M400 224h-24v-72C376 68.2 307.8 0 224 0S72 68.2 72 152v72H48c-26.5 0-48 21.5-48 48v192c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V272c0-26.5-21.5-48-48-48zm-104 0H152v-72c0-39.7 32.3-72 72-72s72 32.3 72 72v72z"></path></svg><!-- <i id="lock" class="fas fa-lock" onclick="editor.editor_mode='fixed'; changeMode('lock');"></i> -->
                <svg id="unlock" class="svg-inline--fa fa-lock-open fa-w-18" onclick="editor.editor_mode='edit'; changeMode('unlock');" style="display: none;" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="lock-open" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 576 512" data-fa-i2svg=""><path fill="currentColor" d="M423.5 0C339.5.3 272 69.5 272 153.5V224H48c-26.5 0-48 21.5-48 48v192c0 26.5 21.5 48 48 48h352c26.5 0 48-21.5 48-48V272c0-26.5-21.5-48-48-48h-48v-71.1c0-39.6 31.7-72.5 71.3-72.9 40-.4 72.7 32.1 72.7 72v80c0 13.3 10.7 24 24 24h32c13.3 0 24-10.7 24-24v-80C576 68 507.5-.3 423.5 0z"></path></svg><!-- <i id="unlock" class="fas fa-lock-open" onclick="editor.editor_mode='edit'; changeMode('unlock');" style="display:none;"></i> -->
              </div>
              <div class="bar-zoom">
                <svg 
                  class="svg-inline--fa fa-search-minus fa-w-16" 
                  onclick="editor.zoom_out()" 
                  aria-hidden="true" 
                  focusable="false" 
                  data-prefix="fas" 
                  data-icon="search-minus" 
                  role="img" 
                  xmlns="http://www.w3.org/2000/svg" 
                  viewBox="0 0 512 512" 
                  data-fa-i2svg="">
                    <path fill="currentColor" 
                      d="M304 192v32c0 6.6-5.4 12-12 12H124c-6.6 0-12-5.4-12-12v-32c0-6.6 5.4-12 12-12h168c6.6 0 12 5.4 12 12zm201 284.7L476.7 505c-9.4 9.4-24.6 9.4-33.9 0L343 405.3c-4.5-4.5-7-10.6-7-17V372c-35.3 27.6-79.7 44-128 44C93.1 416 0 322.9 0 208S93.1 0 208 0s208 93.1 208 208c0 48.3-16.4 92.7-44 128h16.3c6.4 0 12.5 2.5 17 7l99.7 99.7c9.3 9.4 9.3 24.6 0 34zM344 208c0-75.2-60.8-136-136-136S72 132.8 72 208s60.8 136 136 136 136-60.8 136-136z">
                    </path>
                </svg><!-- <i class="fas fa-search-minus" onclick="editor.zoom_out()"></i> -->
                <svg class="svg-inline--fa fa-search fa-w-16" onclick="editor.zoom_reset()" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M505 442.7L405.3 343c-4.5-4.5-10.6-7-17-7H372c27.6-35.3 44-79.7 44-128C416 93.1 322.9 0 208 0S0 93.1 0 208s93.1 208 208 208c48.3 0 92.7-16.4 128-44v16.3c0 6.4 2.5 12.5 7 17l99.7 99.7c9.4 9.4 24.6 9.4 33.9 0l28.3-28.3c9.4-9.4 9.4-24.6.1-34zM208 336c-70.7 0-128-57.2-128-128 0-70.7 57.2-128 128-128 70.7 0 128 57.2 128 128 0 70.7-57.2 128-128 128z"></path></svg><!-- <i class="fas fa-search" onclick="editor.zoom_reset()"></i> -->
                <svg class="svg-inline--fa fa-search-plus fa-w-16" onclick="editor.zoom_in()" aria-hidden="true" focusable="false" data-prefix="fas" data-icon="search-plus" role="img" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512" data-fa-i2svg=""><path fill="currentColor" d="M304 192v32c0 6.6-5.4 12-12 12h-56v56c0 6.6-5.4 12-12 12h-32c-6.6 0-12-5.4-12-12v-56h-56c-6.6 0-12-5.4-12-12v-32c0-6.6 5.4-12 12-12h56v-56c0-6.6 5.4-12 12-12h32c6.6 0 12 5.4 12 12v56h56c6.6 0 12 5.4 12 12zm201 284.7L476.7 505c-9.4 9.4-24.6 9.4-33.9 0L343 405.3c-4.5-4.5-7-10.6-7-17V372c-35.3 27.6-79.7 44-128 44C93.1 416 0 322.9 0 208S93.1 0 208 0s208 93.1 208 208c0 48.3-16.4 92.7-44 128h16.3c6.4 0 12.5 2.5 17 7l99.7 99.7c9.3 9.4 9.3 24.6 0 34zM344 208c0-75.2-60.8-136-136-136S72 132.8 72 208s60.8 136 136 136 136-60.8 136-136z"></path></svg><!-- <i class="fas fa-search-plus" onclick="editor.zoom_in()"></i> -->
              </div>
            </div>
          </div>
        </div>
      </div>
       <div class="col-md-4">
        <div class="row">
          <div class="col-md-12">
            <div class="p-3 mb-5 bg-white rounded" style="height: 30px;">
              <img alt="Vue logo" src="../assets/python_logo.png" style="width: 140px; height: 70px;">
            </div>
            <div class="shadow p-3 mb-5 bg-white rounded" style="margin-top: 50px; height: 300px;">
              # Create a sample collection
users = {'Hans': 'active', 'Éléonore': 'inactive', '景太郎': 'active'}

# Strategy:  Iterate over a copy
for user, status in users.copy().items():
    if status == 'inactive':
        del users[user]

# Strategy:  Create a new collection
active_users = {}
for user, status in users.items():
    if status == 'active':
        active_users[user] = status
            </div>
          </div>
          <div class="p-3 mb-5 bg-white rounded" style="height: 2px;">
            <button type="button" class="btn btn-outline-success" style="margin: 3px;">Run</button>
            <button type="button" class="btn btn-outline-primary">Download</button>
            </div>
        </div>
        <div class="row">
          <div class="col-md-12">
            
            <div class="shadow p-3 mb-5 bg-white rounded" style="margin-top: 50px; height: 300px;">
              result....
              <div class="result-spinner">
                <b-spinner small label="Small Spinner"></b-spinner>
                <b-spinner small label="Small Spinner" type="grow"></b-spinner>
              </div>
            </div>
          </div>
        
        </div>
      </div>
    </div>
</template>

<script>
import Drawflow from 'drawflow';
import Vue from 'vue'

export default {
  data(){
    return {
      editor: null,
      mobile_item_selec: '',
      mobile_last_move: null,
      transform: ''
    }
  },
  mounted: function(){
    this.$nextTick(function(){

      var id = document.getElementById("drawflow");
      this.editor = new Drawflow(id, Vue);
      this.editor.reroute = true;
      this.editor.start();

      var elements = document.getElementsByClassName('drag-drawflow');
      for (var i = 0; i < elements.length; i++) {
        elements[i].addEventListener('touchend', this.drop, false);
        elements[i].addEventListener('touchmove', this.positionMobile, false);
        elements[i].addEventListener('touchstart', this.drag, false );
      }

    })
  },
  methods: {
    positionMobile(ev){
      this.mobile_last_move = ev
    },
    drag(ev){
      if (ev.type === "touchstart") {
        this.mobile_item_selec = ev.target.closest(".drag-drawflow").getAttribute('data-node');
      } else {
      ev.dataTransfer.setData("node", ev.target.getAttribute('data-node'));
      }
    },
    drop(ev){
       if (ev.type === "touchend") {
        var parentdrawflow = document.elementFromPoint( this.mobile_last_move.touches[0].clientX, this.mobile_last_move.touches[0].clientY).closest("#drawflow");
        if(parentdrawflow != null) {
          this.addNodeToDrawFlow(this.mobile_item_selec, this.mobile_last_move.touches[0].clientX, this.mobile_last_move.touches[0].clientY);
        }
        this.mobile_item_selec = '';
      } else {
        ev.preventDefault();
        var data = ev.dataTransfer.getData("node");
        this.addNodeToDrawFlow(data, ev.clientX, ev.clientY);
      }
    },
    allowDrop(ev){
      ev.preventDefault();
    },
    addNodeToDrawFlow(name, pos_x, pos_y){
      if(this.editor.editor_mode === 'fixed') {
        return false;
      }
      pos_x = pos_x * ( this.editor.precanvas.clientWidth / (this.editor.precanvas.clientWidth * this.editor.zoom)) - (this.editor.precanvas.getBoundingClientRect().x * ( this.editor.precanvas.clientWidth / (this.editor.precanvas.clientWidth * this.editor.zoom)));
      pos_y = pos_y * ( this.editor.precanvas.clientHeight / (this.editor.precanvas.clientHeight * this.editor.zoom)) - (this.editor.precanvas.getBoundingClientRect().y * ( this.editor.precanvas.clientHeight / (this.editor.precanvas.clientHeight * this.editor.zoom)));


      switch (name) {
        case 'assignNumber':
        var assignInt = `
        <div>
            <div class="title-box">Assign Number</div>
            <div class="box">
              <p>Variable name</p>
            <input type="text" df-name>
             <p>Value</p>
            <input type="text" df-value>
            </div>
          </div>
        `;
          this.editor.addNode('assignint', 0,  1, pos_x, pos_y, 'assignint', {astType: "Assign", type:"int"}, assignInt );
          break;
        case 'add':
          var add = `
          <div>
            <div class="title-box">Add</div>
          </div>
          `
          this.editor.addNode('add', 2, 1, pos_x, pos_y, 'add', {astType: "Add", op: "Add"}, add );
          break;
        case 'sub':
          var sub = `
          <div>
            <div class="title-box">Sub</div>
          </div>
          `
          this.editor.addNode('sub', 2, 1, pos_x, pos_y, 'sub', {astType: "Sub", op: "Sub"}, sub );
          break;
        case 'mult':
          var mult = `
          <div>
            <div class="title-box">Mult</div>
          </div>
          `
          this.editor.addNode('mult', 2, 1, pos_x, pos_y, 'mult', {astType: "Mult", op: "Mult"}, mult );
          break;
        case 'div':
          var div = `
          <div>
            <div class="title-box">Div</div>
          </div>
          `
          this.editor.addNode('div', 2, 1, pos_x, pos_y, 'div', {astType: "Div", op: "Div"}, div );
          break;
        case 'print':
          var print = `<div>
            <div class="title-box">Print</div>
            <div class="box">
              <input type="text" df-value></div>
            </div>`;
          this.editor.addNode('print', 1, 0, pos_x, pos_y, 'print', {astType: "Print"}, print );
          break;
        case 'ifelse':
          var ifelse = `<div>
                          <div class="title-box">If/Else</div>
                          <div class="box">
                          <p>select operator</p>
                          <Select df-op>
                            <option value="Eq()">Equal</option>
                            <option value="Gt()">Gt</option>
                            <option value="Lt()">Lt</option>
                            <option value="NotEq()">Not Equal</option>
                            <option value="GtE()">GtE</option>
                            <option value="LtE()">LtE()</option>
                          </select>
                            <p>Comparator</p>
                            <input type="text" df-comparators>
                          </div>
                          </div>`;
          this.editor.addNode('ifelse', 1, 2, pos_x, pos_y, 'ifelse', {astType: "IfElse", op : "Eq()"}, ifelse );
          break;
        case 'forIn':
          var forIn = `
          <div>
            <div class="title-box"> For In </div>
            </div>
          </div>
          `;
          this.editor.addNode('forIn', 2, 1, pos_x, pos_y, 'forIn', { astType: "ForIn"}, forIn );
          break;
        case 'code':
         var code = `<div>
            <div class="title-box">Code</div>
            <div class="box">
              <textarea cols="50" rows="2" df-value></textarea>
            </div>`;
          this.editor.addNode('code', 0, 0, pos_x, pos_y, 'code', { astType: "Code"}, code );
          break;
        case 'log':
            var log = `
            <div>
              <div class="title-box"><i class="fas fa-file-signature"></i> Save log file </div>
            </div>
            `;
            this.editor.addNode('log', 1, 0, pos_x, pos_y, 'log', {}, log );
            break;
          case 'google':
            var google = `
            <div>
              <div class="title-box"><i class="fab fa-google-drive"></i> Google Drive save </div>
            </div>
            `;
            this.editor.addNode('google', 1, 0, pos_x, pos_y, 'google', {}, google );
            break;
          case 'email':
            var email = `
            <div>
              <div class="title-box"><i class="fas fa-at"></i> Send Email </div>
            </div>
            `;
            this.editor.addNode('email', 1, 0, pos_x, pos_y, 'email', {}, email );
            break;

          case 'template':
            var template = `
            <div>
              <div class="title-box"><i class="fas fa-code"></i> Template</div>
              <div class="box">
                Ger Vars
                <textarea df-template></textarea>
                Output template with vars
              </div>
            </div>
            `;
            this.editor.addNode('template', 1, 1, pos_x, pos_y, 'template', { "template": 'Write your template'}, template );
            break;
          case 'multiple':
            var multiple = `
            <div>
              <div class="box">
                Multiple!
              </div>
            </div>
            `;
            this.editor.addNode('multiple', 3, 4, pos_x, pos_y, 'multiple', {}, multiple );
            break;
          case 'personalized':
            var personalized = `
            <div>
              Personalized
            </div>
            `;
            this.editor.addNode('personalized', 1, 1, pos_x, pos_y, 'personalized', {}, personalized );
            break;
          case 'dbclick':
            var dbclick = `
            <div>
            <div class="title-box"><i class="fas fa-mouse"></i> Db Click</div>
              <div class="box dbclickbox" ondblclick="showpopup(event)">
                Db Click here
                <div class="modal" style="display:none">
                  <div class="modal-content">
                    <span class="close" onclick="closemodal(event)">&times;</span>
                    Change your variable {name} !
                    <input type="text" df-name>
                  </div>

                </div>
              </div>
            </div>
            `;
            this.editor.addNode('dbclick', 1, 1, pos_x, pos_y, 'dbclick', { name: ''}, dbclick );
            break;

        default:
      }
    },
    showpopup(e){
      e.target.closest(".drawflow-node").style.zIndex = "9999";
      e.target.children[0].style.display = "block";
      //document.getElementById("modalfix").style.display = "block";

      //e.target.children[0].style.transform = 'translate('+translate.x+'px, '+translate.y+'px)';
      this.transform = this.editor.precanvas.style.transform;
      this.editor.precanvas.style.transform = '';
      this.editor.precanvas.style.left = this.editor.canvas_x +'px';
      this.editor.precanvas.style.top = this.editor.canvas_y +'px';
      console.log(this.transform);

      //e.target.children[0].style.top  =  -editor.canvas_y - editor.container.offsetTop +'px';
      //e.target.children[0].style.left  =  -editor.canvas_x  - editor.container.offsetLeft +'px';
      this.editor.editor_mode = "fixed";
      },
      closemodal(e){
        e.target.closest(".drawflow-node").style.zIndex = "2";
        e.target.parentElement.parentElement.style.display  ="none";
        //document.getElementById("modalfix").style.display = "none";
        this.editor.precanvas.style.transform = this.transform;
        this.editor.precanvas.style.left = '0px';
        this.editor.precanvas.style.top = '0px';
        this.editor.editor_mode = "edit";
      },
      changeModule(event){
        var all = document.querySelectorAll(".menu ul li");
        for (var i = 0; i < all.length; i++) {
          all[i].classList.remove('selected');
        }
        event.target.classList.add('selected');
      },
      changeMode(option){
        //console.log(lock.id);
        if(option == 'lock') {
          //lock.style.display = 'none';
          //unlock.style.display = 'block';
        } else {
          //lock.style.display = 'block';
          //unlock.style.display = 'none';
        }
      },
      parse(){
        let data = this.editor.export().drawflow.Home.data;
        let d = [];
        if(data){
          for(var k in data){
            d.push(data[k]);
          }
          this.$store.dispatch('parseCode', d)
        }
      },
      clear(){
        this.editor.clearModuleSelected();
      },
      save(){

      }
  }
}
</script>

<style>
</style>