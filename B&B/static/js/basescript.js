//func to use notie for pop notifications
function notify(msg, type) {
    notie.alert({
        type: type,
        text: msg,
    })
}

//boilerplate funcs to use SweetAlert for popups
function Prompt() {
    const toast = ({ msg = "", icon = "success", position = "top-end" }) => {
        const Toast = Swal.mixin({
            toast: true,
            title: msg,
            position: position,
            icon: icon,
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
                toast.addEventListener('mouseenter', Swal.stopTimer)
                toast.addEventListener('mouseleave', Swal.resumeTimer)
            }
        })
        Toast.fire({})
    }
    const success = ({ title = "", text = "", icon, footer = "" }) => {
        Swal.fire({
            title: title,
            text: text,
            icon: 'success',
            footer: footer,
        })
    }
    const error = ({ title = "", text = "", icon, footer = "" }) => {
        Swal.fire({
            title: title,
            text: text,
            icon: 'error',
            footer: footer,
        })
    }
    const custom = async (c) => {
        const { title = "", text = "", icon="", showConfirmButton=true } = c;
        const { value: formValues } = await Swal.fire({
            title: title,
            html: text,
            icon:icon,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            showConfirmButton:showConfirmButton,
            willOpen: () => {
                if (c.willOpen !== undefined) {
                    c.willOpen();
                }
            },
            didOpen: () => {
                if (c.didOpen !== undefined) {
                    c.didOpen();
                }
            }
        })

        if (formValues && formValues.dismiss !== Swal.DismissReason.cancel && formValues.value !== "") {
            if (c.callback !== undefined) {
                c.callback(formValues);
            } else {
                c.callback(false);
            }
        }
    }
    return {
        toast: toast,
        success: success,
        error: error,
        custom: custom
    }
}
let attention = Prompt();