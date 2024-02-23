const viewPage = document.getElementById("view-page");
const editPage = document.getElementById("edit-page");

const editToggleBtn = document.querySelectorAll(".edit-toggle");

for (const button of editToggleBtn) {
  button.addEventListener("click", () => {
    viewPage.classList.toggle("hidden");
    editPage.classList.toggle("hidden");
  });
}
