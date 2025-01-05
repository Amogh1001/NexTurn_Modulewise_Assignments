const amount = document.getElementById("amount");
const description = document.getElementById("description");
const category = document.getElementById("category");
const listContainer = document.getElementById("listContainer");

function addExpense() {
    if (amount.value === "" || description.value === "" || category.value === "") {
        alert("Please fill in all fields before adding an expense.");
    } else {
        const li = document.createElement("li");
        li.draggable = true; // Make the list item draggable
        listContainer.appendChild(li);

        // Create content elements
        const expenseDetails = document.createElement("p");
        expenseDetails.innerHTML = `<strong>Amount:</strong> $${amount.value} | <strong>Description:</strong> ${description.value} | <strong>Category:</strong> ${category.value}`;

        // Create buttons
        const editButton = document.createElement("button");
        const deleteButton = document.createElement("button");
        editButton.innerHTML = "Edit";
        deleteButton.innerHTML = "Delete";

        // Append to list item
        li.appendChild(expenseDetails);
        li.appendChild(deleteButton);
        li.appendChild(editButton);
    }

    // Clear input fields
    amount.value = "";
    description.value = "";
    category.value = "Food";

    storeData();
}

// Event listeners for editing and deleting
listContainer.addEventListener("click", (e) => {
    if (e.target.tagName === "BUTTON" && e.target.innerHTML === "Delete") {
        e.target.parentElement.remove();
        storeData();
    } else if (e.target.tagName === "BUTTON" && e.target.innerHTML === "Edit") {
        const li = e.target.parentElement;
        const expenseDetails = li.firstChild.textContent.match(/Amount: \$(\d+(\.\d{1,2})?) \| Description: (.+) \| Category: (.+)/);

        // Extract details
        const newAmount = prompt("Enter new amount", expenseDetails[1]);
        const newDescription = prompt("Enter new description", expenseDetails[3]);
        const newCategory = prompt("Enter new category", expenseDetails[4]);

        // Update details
        if (newAmount && newDescription && newCategory) {
            li.firstChild.innerHTML = `<strong>Amount:</strong> $${newAmount} | <strong>Description:</strong> ${newDescription} | <strong>Category:</strong> ${newCategory}`;
            storeData();
        }
    }
});

function storeData() {
    localStorage.setItem("expenses", listContainer.innerHTML);
}

function showData() {
    const storedData = localStorage.getItem("expenses");
    if (storedData) {
        listContainer.innerHTML = storedData;
        addDragAndDropHandlers(); // Re-add drag handlers
    }
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

// Initialize
showData();
addDragAndDropHandlers();
