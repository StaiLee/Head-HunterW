import conversations from "./conversations.js";

var API_URL = "http://localhost:1507"

var timeouts = []

const makeMessageTimeout = (i, selector, message) => {
    timeouts.push(
        setTimeout(() => {
            document.querySelector(selector).innerHTML += message[i]
        }, i * 50)
    )
}

const clearAllTimeouts = () => timeouts.forEach(t => clearTimeout(t))

const animateMessage = async (selector, message) => {
    for (var i = 0; i < message.length; ++i) {
        makeMessageTimeout(i, selector, message)
    }
}

const sendRequest = (path, body) => {
    var method = "GET"
    if (body) method = "POST"
    fetch(API_URL + path, {
        body,
        method
    })
}

const formPop = async () => {
    var search = document.getElementById('search');
    var button = document.getElementById('button');
    var input = document.getElementById('input');

    function loading() {
        document.getElementById('search').classList.add('loading');

        setTimeout(function () {
            document.getElementById('search').classList.remove('loading');
        }, 1500);
    }

    button.addEventListener('click', loading);

    input.addEventListener('keydown', function () {
        // @ts-ignore
        if (event.keyCode == 13) loading(), sendRequest("/api", `{'message': '${document.querySelector('#input').value}'}`);
    });
}


// @ts-ignore
const showParticle = () => particlesJS('particles-js', {
    "particles": {
        "number": {
            "value": 50,
            "density": {
                "enable": true,
                "value_area": 800
            }
        },
        "color": {
            "value": "#ffffff"
        },
        "shape": {
            "type": "circle",
            "stroke": {
                "width": 5,
                "color": "#de5fb1"
            },
            "polygon": {
                "nb_sides": 5
            },
            "image": {
                "src": "img/github.svg",
                "width": 100,
                "height": 100
            }
        },
        "opacity": {
            "value": 0.5,
            "random": false,
            "anim": {
                "enable": false,
                "speed": 1,
                "opacity_min": 0.1,
                "sync": false
            }
        },
        "size": {
            "value": 5,
            "random": true,
            "anim": {
                "enable": false,
                "speed": 40,
                "size_min": 0.1,
                "sync": false
            }
        },
        "line_linked": {
            "enable": true,
            "distance": 150,
            "color": "#ffffff",
            "opacity": 0.4,
            "width": 1
        },
        "move": {
            "enable": true,
            "speed": 5,
            "direction": "none",
            "random": false,
            "straight": false,
            "out_mode": "out",
            "attract": {
                "enable": false,
                "rotateX": 600,
                "rotateY": 1200
            }
        }
    },
    "interactivity": {
        "detect_on": "canvas",
        "events": {
            "onhover": {
                "enable": true,
                "mode": "repulse"
            },
            "onclick": {
                "enable": true,
                "mode": "push"
            },
            "resize": true
        },
        "modes": {
            "grab": {
                "distance": 400,
                "line_linked": {
                    "opacity": 1
                }
            },
            "bubble": {
                "distance": 400,
                "size": 40,
                "duration": 2,
                "opacity": 8,
                "speed": 3
            },
            "repulse": {
                "distance": 200
            },
            "push": {
                "particles_nb": 4
            },
            "remove": {
                "particles_nb": 2
            }
        }
    },
    "retina_detect": true,
    "config_demo": {
        "hide_card": false,
        "background_color": "#663399",
        "background_image": "",
        "background_position": "50% 50%",
        "background_repeat": "no-repeat",
        "background_size": "cover"
    }
});

document.addEventListener("DOMContentLoaded", async () => {
    var conv_index = 0;
    var display = true;
    document.body.innerHTML += conversations.conversation_div_start

    await animateMessage(".npc_message", conversations.woman[conv_index++] + " [ENTER]")
    document.addEventListener("keydown", async (key) => {
        var conv_message = document.querySelector(".npc_message")
        if (key.keyCode != 13 || !conv_message) return
        clearAllTimeouts()
        conv_message.innerHTML = ""
        if (conv_index < 0) conv_index = 0
        else if (conv_index == conversations.woman.length) {
            display = false;
            console.log("HERE", display)
            document.body.innerHTML += conversations.conversation_form
            document.body.innerHTML += conversations.conversation_form_css
            document.querySelector(".npc_message").innerHTML = ""
            conv_index++
            showParticle()
            await formPop()
        }
        if (display) await animateMessage(".npc_message", conversations.woman[conv_index++] + " [ENTER]")
    })

    showParticle()
})