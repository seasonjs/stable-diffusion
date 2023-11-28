# Model Convert Script

## Requirements

- vocab.json, from https://huggingface.co/openai/clip-vit-large-patch14/raw/main/vocab.json


```shell
pip install -r requirements.txt
```

## Usage
```
usage: convert.py [-h] [--out_type {f32,f16,q4_0,q4_1,q5_0,q5_1,q8_0}] [--out_file OUT_FILE] model_path

Convert Stable Diffuison model to GGML compatible file format

positional arguments:
  model_path            model file path (*.pth, *.pt, *.ckpt, *.safetensors)

options:
  -h, --help            show this help message and exit
  --out_type {f32,f16,q4_0,q4_1,q5_0,q5_1,q8_0}
                        output format (default: based on input)
  --out_file OUT_FILE   path to write to; default: based on input and current working directory
```

## Usage
```
usage: convert.exe [MODEL_PATH] --type [OUT_TYPE] [arguments]
Model supported for conversion: .safetensors models or .ckpt checkpoints models

arguments:
  -h, --help                         show this help message and exit
  -o, --out [FILENAME]               path or name to converted model
  --vocab [FILENAME]                 path to custom vocab.json (usually unnecessary)
  -v, --verbose                      print processing info - dev info
  -l, --lora                         force read the model as a LoRA
  --vae [FILENAME]                   merge a custom VAE
  -t, --type [OUT_TYPE]              output format (f32, f16, q4_0, q4_1, q5_0, q5_1, q8_0)
```