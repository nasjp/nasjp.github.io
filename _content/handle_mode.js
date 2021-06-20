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

const init = (checkbox) => {
    checkbox.addEventListener("change", () => {
        if (checkbox.checked) {
            setDark(checkbox);
            return;
        }

        setNormal(checkbox);
    });

    if (sessionStorage.getItem("mode") === "dark") {
        setDark(checkbox);
        return;
    }

    setNormal(checkbox);
};

init(document.getElementById("handle-mode"));