package provider

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLocalFile() *schema.Resource {

	return &schema.Resource{

		Description: "This resource will support create, read, update and delete file resource via Terraform.",

		CreateContext: resourceCreateFile,
		ReadContext:   resourceReadFile,
		UpdateContext: resourceUpdateFile,
		DeleteContext: resourceDeleteFile,

		Schema: map[string]*schema.Schema{

			"file_name": {
				Description: "Name of the file.",
				Type:        schema.TypeString,
				Required:    true,
			},

			"file_content": {
				Description: "Content that must be stored/retrieve to/from the file.",
				Type:        schema.TypeString,
				Required:    true,
			},
		},
	}
}

func resourceCreateFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {

	//Retrieve the inputs user provided in the terrform .tf script file
	filename := d.Get("file_name").(string)
	content := d.Get("file_content").(string)

	myfile, err := os.Create(filename)

	if err != nil {
		panic(err)
	}

	myfile.WriteString(content)
	myfile.Sync() //This will ensure the file content is flushed out (written) to disk

	d.SetId("resource-101")

	resourceReadFile(ctx, d, meta)

	return nil
}

func resourceReadFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve the inputs user provided in the terrform .tf script file
	filename := d.Get("file_name").(string)

	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	d.Set("content", content)

	return nil
}

func resourceUpdateFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	return nil
}

func resourceDeleteFile(ctx context.Context, d *schema.ResourceData, meta any) diag.Diagnostics {
	//Retrieve the inputs user provided in the terrform .tf script file
	filename := d.Get("file_name").(string)

	err := os.Remove(filename)

	if err != nil {
           log.Fatal(err)
	}

	return nil
}
