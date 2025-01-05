const listItem = document.getElementById("listItem");
const listContainer = document.getElementById("listContainer");

function addItem() {
    if (listItem.value === "") {
        alert("Please enter text to add to the list");
    } else {
        const li = document.createElement("li");
        li.draggable = true; // Make the list item draggable
        listContainer.appendChild(li);
        const content = document.createElement("h3");
        content.innerHTML = listItem.value;
        const editButton = document.createElement("button");
        const deleteButton = document.createElement("button");
        editButton.innerHTML = "Edit";
        deleteButton.innerHTML = "Delete";
        li.appendChild(content);
        li.appendChild(deleteButton);
        li.appendChild(editButton);
    }
    listItem.value = "";
    storeData();
}

listContainer.addEventListener("click", (e) => {
    if (e.target.tagName === "LI") {
        if (e.target.style.textDecoration === "line-through") {
            e.target.style.textDecoration = "none";
            storeData();
        } else {
            e.target.style.textDecoration = "line-through";
            storeData();
        }
    } else if (e.target.tagName == "BUTTON" && e.target.innerHTML === "Delete") {
        e.target.parentElement.remove();
        storeData();
    } else if (e.target.tagName == "BUTTON" && e.target.innerHTML === "Edit") {
        const newText = prompt("Enter new text", e.target.parentElement.firstChild.textContent);
        e.target.parentElement.firstChild.textContent = newText;
        storeData();
    }
});

function storeData() {
    localStorage.setItem("data", listContainer.innerHTML);
}

function showData() {
    listContainer.innerHTML = localStorage.getItem("data");
    addDragAndDropHandlers(); // Re-add drag-and-drop handlers after restoring data
}

function clearData() {
    localStorage.clear();
    listContainer.innerHTML = "";
}

// Drag and Drop Handlers
function addDragAndDropHandlers() {
    const items = listContainer.querySelectorAll("li");
    items.forEach((item) => {
        item.draggable = true; // Ensure items are draggable

        item.addEventListener("dragstart", (e) => {
            e.dataTransfer.setData("text/plain", e.target.id);
            e.target.classList.add("dragging");
        });

        item.addEventListener("dragend", (e) => {
            e.target.classList.remove("dragging");
        });
    });

    listContainer.addEventListener("dragover", (e) => {
        e.preventDefault();
        const draggingItem = document.querySelector(".dragging");
        const closestItem = getClosestItem(listContainer, e.clientY);
        if (closestItem == null) {
            listContainer.appendChild(draggingItem);
        } else {
            listContainer.insertBefore(draggingItem, closestItem);
        }
    });

    listContainer.addEventListener("drop", () => {
        storeData(); // Save the updated order to localStorage
    });
}

// Helper function to find the closest item
function getClosestItem(container, y) {
    const items = [...container.querySelectorAll("li:not(.dragging)")];
    return items.reduce((closest, child) => {
        const box = child.getBoundingClientRect();
        const offset = y - box.top - box.height / 2;
        if (offset < 0 && offset > closest.offset) {
            return { offset, element: child };
        } else {
            return closest;
        }
    }, { offset: Number.NEGATIVE_INFINITY }).element;
}

showData();
addDragAndDropHandlers();
