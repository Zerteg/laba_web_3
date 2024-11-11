  //пытался сделать всплывающее окошко регистрации 


document.addEventListener("DOMContentLoaded", function() {

  // Модальное окно регистрации
    const registerModal = document.getElementById("Register");
    const registerBtn = document.getElementById("registerBtn");
    const registerCloseBtn = registerModal.querySelector(".close");
  
    // Открытие модального окна регистрации
    registerBtn.onclick = function() {
      registerModal.style.display = "flex";
    };
  
    // Закрытие модального окна регистрации при нажатии на кнопку закрытия
    registerCloseBtn.onclick = function() {
      registerModal.style.display = "none";
    };
  
    // Закрытие окна при клике вне области окна для регистрации
    window.onclick = function(event) {
      if (event.target.closest(".modal-content") === null && event.target === registerModal) {
        registerModal.style.display = "none";
      }
    };
  
    // Закрытие модального окна регистрации при нажатии клавиши Escape
    document.addEventListener("keydown", function(event) {
      if (event.key === "Escape") {
        registerModal.style.display = "none";
      }
    });
  
    // Модальное окно обратной связи
    const feedbackModal = document.getElementById("myModal");
    const feedbackBtn = document.getElementById("feedback");
    const feedbackCloseBtn = feedbackModal.querySelector(".close");
  
    // Открытие модального окна обратной связи
    feedbackBtn.onclick = function() {
      feedbackModal.style.display = "flex";
    };
  
    // Закрытие модального окна обратной связи при нажатии на кнопку закрытия
    feedbackCloseBtn.onclick = function() {
      feedbackModal.style.display = "none";
    };
  
    // Закрытие окна при клике вне области окна для обратной связи
    window.onclick = function(event) {
      if (event.target.closest(".modal-content") === null && event.target === feedbackModal) {
        feedbackModal.style.display = "none";
      }
    };
  
    // Закрытие модального окна обратной связи при нажатии клавиши Escape
    document.addEventListener("keydown", function(event) {
      if (event.key === "Escape") {
        feedbackModal.style.display = "none";
      }
    });
  });
  