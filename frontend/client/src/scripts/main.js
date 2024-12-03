

function login() {
  const email = document.getElementById('login-email').value;
  const password = document.getElementById('login-password').value;


  if (email === "" || password === "") {
    alert('Please enter a valid email password');
    return
  }

  window.location.href = "../pages/home.html";
}

function register() {
  const email = document.getElementById('register-email').value;
  const password = document.getElementById('register-password').value;
  const confirmPassword = document.getElementById('register-confirm-password').value;

  if (password === confirmPassword) {
    alert(`Register with Email: ${email}`);
  } else {
    alert('Passwords do not match!');
  }
}

function searchVideos(event) {
  // Prevent form submission (page reload)
  event.preventDefault();

  // Log for debugging
  console.log('Search Videos');

  // Data array for video items
  const videoItems = [
    {
      href: "video.html",
      imgSrc: "https://avatars.mds.yandex.net/i?id=2f203721ce56392856c96a2387c935f5_l-12752373-images-thumbs&n=13",
      title: "Video 1"
    },
    {
      href: "video.html",
      imgSrc: "https://i.pinimg.com/originals/8e/43/05/8e4305e5ca2524e022a75c5fdf0f1803.jpg",
      title: "video 2"
    },
    {
      href: "login.html",
      imgSrc: "https://i.pinimg.com/originals/f9/56/ac/f956ac4e28a0f7c30c76dc7647983b8a.jpg",
      title: "video 3"
    }
  ];

  // Select the container element
  const videoContainer = document.getElementById('videoContainer');
  // Clear previous search results before adding new ones
  videoContainer.innerHTML = '';

  // Function to dynamically create video items
  videoItems.forEach(item => {
    // Create the main video-item div
    const videoItemDiv = document.createElement('div');
    videoItemDiv.className = 'video-item';

    // Create the anchor element
    const anchor = document.createElement('a');
    anchor.href = item.href;

    // Create the image element
    const img = document.createElement('img');
    img.src = item.imgSrc;
    img.alt = 'video';
    img.className = 'rounded mx-auto d-block';
    img.width = 200;
    img.height = 200;

    // Create the title element
    const title = document.createElement('h5');
    title.textContent = item.title;

    // Append image and title to the anchor
    anchor.appendChild(img);
    anchor.appendChild(title);

    // Append anchor to the main div
    videoItemDiv.appendChild(anchor);

    // Append the video-item div to the container
    videoContainer.appendChild(videoItemDiv);
  });
}

const searchForm = document.getElementById('searchForm');
searchForm.addEventListener('submit', searchVideos);
