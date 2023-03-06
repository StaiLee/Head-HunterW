export default {

  woman: [
    "Hi Hacker, welcome to the lounge !",
    "On which head would you like to have some infos ?",
    "Enter his name in the input, what are you waiting for ?"
  ],
  hacker_man_left: [

  ],
  hacker_man_right_1: [

  ],
  hacker_man_right_2: [

  ],


  conversation_form: ` 
    <div id="search">
    <input id="input" placeholder="Type your head..."/>
    <button id="button"><i class="fa fa-search"></i></button>
    <div class="spinner"><i class="fa fa-spinner"></i></div>
  </div>
  <div class="note">Click the button or hit enter.</div>
    `,

  conversation_form_css: `
    <style>
    .note {
        font-size: 0.375em;
        font-weight: bold;
        text-transform: uppercase;
        color: #6426a7;
      }
      
      #search {
        position: absolute;
        align-items: center;
        background: #8a44d5;
        border-radius: 10px;
        display: flex;
        justify-content: space-between;
        margin: 0.5em 0;
        padding: 0.5em 0.5em 0.5em 1em;
        transition: all 0.5s;
        width: 500px;
      }
      #search:hover, #search:focus {
        background: #853cd3;
      }
      #search button,
      #search input {
        -webkit-appearance: none;
           -moz-appearance: none;
                appearance: none;
        background: transparent;
        border: 0;
        color: inherit;
        font: inherit;
        outline: 0;
      }
      #search button {
        cursor: pointer;
        padding: 0 0.25em;
      }
      #search input {
        flex: 1;
      }
      #search input::-moz-placeholder {
        color: #fff;
      }
      #search input:-ms-input-placeholder {
        color: #fff;
      }
      #search input::placeholder {
        color: #fff;
      }
      #search .spinner {
        -webkit-animation: spinner 1s infinite linear;
                animation: spinner 1s infinite linear;
        display: none;
        padding: 0 0.25em;
      }
      
      #search.loading button {
        display: none;
      }
      #search.loading .spinner {
        display: block;
      }
      
      @-webkit-keyframes spinner {
        0% {
          transform: rotate(0deg);
        }
        100% {
          transform: rotate(360deg);
        }
      }
      
      @keyframes spinner {
            0% {
            transform: rotate(0deg);
            }
            100% {
            transform: rotate(360deg);
            }
        }
    </style>
    `,

  conversation_div_start: `
    <div class="npc_conversation" style="position: absolute; width: 100vh; z-index: 50;">
        <p class="npc_message" style="color:white; font-size: 50px;"></p>
    </div>
`
}