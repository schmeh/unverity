
document.addEventListener('DOMContentLoaded', async () => {
    document.getElementById('start').addEventListener("click", getDissections)
});

function getDissections() {
    const soloCallout = document.getElementById('soloCallout').value;
    const shape1 = document.getElementById('object1').value;
    const shape2 = document.getElementById('object2').value;
    const shape3 = document.getElementById('object3').value;
    const dissections = document.getElementById('dissections');
    dissections.innerText = dissect(soloCallout, shape1, shape2, shape3)
}
