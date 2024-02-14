const music = new Audio('vande.mp3');

const songs = [{
        id: '1',
        songName: `On My Way`,
        subtitle: `Alan Walker`,
        poster: "1.jpg"
    },
    // ... other song objects ...
];

// Update song list
Array.from(document.getElementsByClassName('songItem')).forEach((element, i) => {
    if (i < songs.length) {
        const imgElement = element.getElementsByTagName('img')[0];
        const titleElement = element.getElementsByTagName('h5')[0];

        if (imgElement) imgElement.src = songs[i].poster;
        if (titleElement) {
            titleElement.innerHTML = `${songs[i].songName} <br><div class="subtitle">${songs[i].subtitle}</div>`;
        }
    }
});

// Play/Pause functionality
let masterPlay = document.getElementById('masterPlay');
let wave = document.getElementsByClassName('wave')[0];
if (masterPlay) {
    masterPlay.addEventListener('click', () => {
        if (music.paused || music.currentTime <= 0) {
            music.play();
            masterPlay.classList.remove('bi-play-fill');
            masterPlay.classList.add('bi-pause-fill');
            wave.classList.add('active2');
        } else {
            music.pause();
            masterPlay.classList.add('bi-play-fill');
            masterPlay.classList.remove('bi-pause-fill');
            wave.classList.remove('active2');
        }
    });
}

// Additional functionality to play a specific song when clicked
const playListPlays = document.getElementsByClassName('playListPlay');
Array.from(playListPlays).forEach((playIcon, index) => {
    playIcon.addEventListener('click', () => {
        if (index < songs.length) {
            music.src = `path/to/your/songs/${songs[index].id}.mp3`; // Update the path as needed
            music.play();
            // Update UI accordingly

            // Update the currently playing song index
            aceille = index;

            // Update the icon classes
            Array.from(document.getElementsByClassName('playListPlay')).forEach((element, i) => {
                element.classList.remove('bi-pause-circle-fill');
                if (i === index) {
                    element.classList.add('bi-pause-circle-fill');
                } else {
                    element.classList.add('bi-play-circle-fill');
                }
            });
        }
    });
});

let aceille = 0; // Initialize aceille with 0

Array.from(document.getElementsByClassName('playListPlay')).forEach((element, index) => {
    element.addEventListener('click', () => {
        aceille = index; // Update aceille with the index of the clicked song
    });
});