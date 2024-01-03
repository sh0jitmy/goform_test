const generateForm = async () => {
  try {
    const formData = await new Promise((resolve, reject) => {
      // ここでWasmからのデータを取得する等、非同期処理を行う
      const result = generateFormData(); // ここで適切なデータを取得する必要があります
      if (result) {
        resolve(result);
      } else {
        reject(new Error('Failed to generate form data'));
      }
    });

    const formContainer = document.getElementById('form-container');

    formData.forEach(field => {
      const label = document.createElement('label');
      label.setAttribute('for', field.name);
      label.innerText = field.label;

      const input = document.createElement(field.input_type === 'textarea' ? 'textarea' : 'input');
      input.setAttribute('type', field.input_type);
      input.setAttribute('id', field.name);
      input.setAttribute('name', field.name);

      formContainer.appendChild(label);
      formContainer.appendChild(input);
      formContainer.appendChild(document.createElement('br'));
      formContainer.appendChild(document.createElement('br'));
    });
  } catch (error) {
    console.error(error);
    // エラーが発生した場合の処理を記述する（例: エラーメッセージを表示するなど）
  }
};

generateForm();

