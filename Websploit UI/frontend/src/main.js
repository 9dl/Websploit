const totalVisitedElement = document.getElementById('totalVisited');
const duplicatedElement = document.getElementById('duplicated');
const todayVisitedElement = document.getElementById('todayVisited');
const mostUsedElement = document.getElementById('mostUsed');
const blacklistedElement = document.getElementById('blacklisted');
const urlsElement = document.getElementById('Urls');

function UpdateTotalVisited(number) {
    totalVisitedElement.textContent = number;
}

function UpdateDuplicated(number) {
    duplicatedElement.textContent = number;
}

function UpdateTodayVisited(number) {
    todayVisitedElement.textContent = number;
}

function UpdateMostUsed(number) {
    mostUsedElement.textContent = number;
}

function UpdateBlacklisted(number) {
    blacklistedElement.textContent = number;
}

function UpdateUrls(urls) {
    urlsElement.innerHTML += "<br>"+urls;
}
