
document.addEventListener('DOMContentLoaded', function() {

    const avatar = document.getElementById('avatar');
    const slideOutMenu = document.getElementById('slideOutMenu');

    avatar.addEventListener('click', function() {
        slideOutMenu.classList.toggle('active');
    });

    document.addEventListener('click', function(event) {
        if (!avatar.contains(event.target) && !slideOutMenu.contains(event.target)) {
            slideOutMenu.classList.remove('active');
        }
    });
    
});
