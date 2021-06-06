function notify(msg, type) {
    notie.alert({
        type: type,
        text: msg,
    })
}
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
        const { title = "", text = "" } = c;
        const { value: formValues } = await Swal.fire({
            title: title,
            html: text,
            backdrop: false,
            focusConfirm: false,
            showCancelButton: true,
            willOpen: () => {
                const elem = document.getElementById("reservation-date-modal");
                const rp = new DateRangePicker(elem, {
                    format: 'yyyy-mm-dd',
                    showOnFocus: true,
                })
            },
            preConfirm: () => {
                return [
                    document.getElementById('start_date').value,
                    document.getElementById('end_date').value
                ]
            },
            didOpen: () => {
                document.getElementById('start_date').removeAttribute('disabled');
                document.getElementById('end_date').removeAttribute('disabled');
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