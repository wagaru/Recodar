import { useState } from "react"
import { useForm } from "react-hook-form"
import TextField from '@material-ui/core/TextField'
import Button from '@material-ui/core/Button'
import DialogTitle from '@material-ui/core/DialogTitle'
import Dialog from '@material-ui/core/Dialog'
import DialogActions from '@material-ui/core/DialogActions'
import DialogContent from '@material-ui/core/DialogContent'
import DialogContentText from '@material-ui/core/DialogContentText'

export default function VideoPage() {
  const {handleSubmit, register} = useForm()
  const [open, setOpen] = useState(false)

  const openDialog = () => {
    setOpen(true)
  }

  const closeDialog = () => {
    setOpen(false)
  }

  const onSubmit = data => {
    console.log(data)
    closeDialog()
  }

  return (
    <div>
      <Button variant="outlined" color="primary" onClick={openDialog}>Upload</Button>
      <Dialog open={open} onClose={closeDialog} aria-labelledby="form-dialog-title">
        <DialogTitle id="form-dialog-title">上傳影片</DialogTitle>
        <DialogContent>
          <form onSubmit={handleSubmit(onSubmit)}>
          <DialogContentText>
            上傳影片前請再三確認不會洩漏你不想要的資訊.
          </DialogContentText>
            <TextField
              required
              autoFocus
              margin="dense"
              name="url"
              label="Youtube 影片網址"
              type="url"
              fullWidth
              inputRef={register}
            />
            <TextField
              required
              autoFocus
              margin="dense"
              name="city"
              label="縣/市"
              type="text"
              fullWidth
              inputRef={register}
            />
            <TextField
              required
              autoFocus
              margin="dense"
              name="road"
              label="道路"
              type="text"
              fullWidth
              inputRef={register}
            />
            <TextField
              required
              autoFocus
              margin="dense"
              name="time"
              label="時間"
              type="text"
              fullWidth
              inputRef={register}
            />
            <TextField
              required
              autoFocus
              margin="dense"
              name="info"
              label="描述"
              type="text"
              fullWidth
              inputRef={register}
            />
          <DialogActions>
            <Button onClick={closeDialog} color="primary">
              取消
            </Button>
            <Button type="submit" color="primary">
              上傳
            </Button>
          </DialogActions>
          </form>
        </DialogContent>
      </Dialog>
    </div>
  )
}