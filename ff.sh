
#!/bin/bash
curl -X POST \
-F 'text="новый релиз"' \
-F 'reply_markup={"inline_keyboard":[[{"text":"Deploy","callback_data":"{\"IsBot\":true}"}]]}' \
https://api.telegram.org/bot8129437003:AAEoEFUfRxgJIAnvz5KkolOtohPyUjoUn_0/sendMessage?chat_id=100130785