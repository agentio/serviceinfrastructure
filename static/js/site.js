
// https://www.dannyguo.com/blog/how-to-add-copy-to-clipboard-buttons-to-code-blocks-in-hugo/

if (navigator && navigator.clipboard) {
    addCopyButtons(navigator.clipboard);
} else {
    var script = document.createElement('script');
    script.src = 'https://cdnjs.cloudflare.com/ajax/libs/clipboard-polyfill/2.7.0/clipboard-polyfill.promise.js';
    script.integrity = 'sha256-waClS2re9NUbXRsryKoof+F9qc1gjjIhc2eT7ZbIv94=';
    script.crossOrigin = 'anonymous';
    script.onload = function () {
        addCopyButtons(clipboard);
    };

    document.body.appendChild(script);
}

function addCopyButtons(clipboard) {
    document.querySelectorAll('pre > code').forEach(function (codeBlock) {
        var pre = codeBlock.parentNode;
        if (codeBlock.classList.contains('language-prompt')) {
            var button = document.createElement('button');
            button.className = 'copy-code-button';
            button.type = 'button';
            button.innerText = 'copy↯';
            button.addEventListener('click', function () {
                clipboard.writeText(codeBlock.innerText).then(function () {
                    /* Chrome doesn't seem to blur automatically,
                       leaving the button in a focused state. */
                    button.blur();
                    button.innerText = 'copied!';
                    setTimeout(function () {
                        button.innerText = 'copy↯';
                    }, 2000);
                }, function (error) {
                    button.innerText = 'Error';
                });
            });
            pre.parentNode.insertBefore(button, pre);
            pre.style.border = "1px dashed";
        }
    });
}
