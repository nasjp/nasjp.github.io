const init = (checkbox) => {
    checkbox.addEventListener("change", function() {
        if (checkbox.checked) {
            setDark(checkbox);
            return;
        }

        setNormal(checkbox);
    });

    if (sessionStorage.getItem("mode") == "dark") {
        setDark(checkbox);
        return;
    }

    setNormal(checkbox);
};

const setDark = (checkbox) => {
    document.body.classList.add("dark-mode");
    checkbox.checked = true;
    sessionStorage.setItem("mode", "dark");
};

const setNormal = (checkbox) => {
    document.body.classList.remove("dark-mode");
    checkbox.checked = false;
    sessionStorage.setItem("mode", "light");
};

init(document.getElementById("handle-mode"));