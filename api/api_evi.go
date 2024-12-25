package api

import "fmt"

// EviGetTools returns a list of tools
// https://dev.hume.ai/reference/empathic-voice-interface-evi/tools/list-tools
func (h *HumeApi) EviGetTools(pageNumber int64, pageSize int64, restrictToMostRecent bool, name string) (*ToolsList, error) {
	h.printDebug("GetTools")

	var toolsList ToolsList
	var err error

	if pageNumber < 0 {
		pageNumber = 0
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	toolsList.PageNumber = pageNumber
	toolsList.PageSize = pageSize

	url := fmt.Sprintf("%s/tools?page_number=%d&page_size=%d&restrict_to_most_recent=%t&name=%s", h.BaseUrl, pageNumber, pageSize, restrictToMostRecent, name)
	h.printDebug(url)

	err = h.get(url, &toolsList)
	if err != nil {
		return nil, err
	}

	return &toolsList, nil
}

// List tool versions
//GET
//https://api.hume.ai/v0/evi/tools/:id
//Fetches a list of a Tool’s versions.
//
//Refer to our tool use guide for comprehensive instructions on defining and integrating tools into EVI.
//
//Path parameters
//id
//string
//Required
//Identifier for a Tool. Formatted as a UUID.
//
//Query parameters
//page_number
//integer
//Optional
//Specifies the page number to retrieve, enabling pagination.
//
//This parameter uses zero-based indexing. For example, setting page_number to 0 retrieves the first page of results (items 0-9 if page_size is 10), setting page_number to 1 retrieves the second page (items 10-19), and so on. Defaults to 0, which retrieves the first page.
//
//page_size
//integer
//Optional
//Specifies the maximum number of results to include per page, enabling pagination. The value must be between 1 and 100, inclusive.
//
//For example, if page_size is set to 10, each page will include up to 10 items. Defaults to 10.
//
//restrict_to_most_recent
//boolean
//Optional
//Defaults to false
//By default, restrict_to_most_recent is set to true, returning only the latest version of each tool. To include all versions of each tool in the list, set restrict_to_most_recent to false.
//
//Response
//Success
//
//page_number
//integer
//The page number of the returned list.
//
//This value corresponds to the page_number parameter specified in the request. Pagination uses zero-based indexing.
//
//page_size
//integer
//The maximum number of items returned per page.
//
//This value corresponds to the page_size parameter specified in the request.
//
//total_pages
//integer
//The total number of pages in the collection.
//
//tools_page
//list of optional objects
//List of tools returned for the specified page_number and page_size.
//
//
//Hide 11 properties
//tool_type
//"BUILTIN" or "FUNCTION"
//Allowed values:
//BUILTIN
//FUNCTION
//Type of Tool. Either BUILTIN for natively implemented tools, like web search, or FUNCTION for user-defined tools.
//
//id
//string
//Identifier for a Tool. Formatted as a UUID.
//
//version
//integer
//Version number for a Tool.
//
//Tools, Configs, Custom Voices, and Prompts are versioned. This versioning system supports iterative development, allowing you to progressively refine tools and revert to previous versions if needed.
//
//Version numbers are integer values representing different iterations of the Tool. Each update to the Tool increments its version number.
//
//version_type
//"FIXED" or "LATEST"
//Allowed values:
//FIXED
//LATEST
//Versioning method for a Tool. Either FIXED for using a fixed version number or LATEST for auto-updating to the latest version.
//
//name
//string
//Name applied to all versions of a particular Tool.
//
//created_on
//long
//Time at which the Tool was created. Measured in seconds since the Unix epoch.
//
//modified_on
//long
//Time at which the Tool was last modified. Measured in seconds since the Unix epoch.
//
//parameters
//string
//Stringified JSON defining the parameters used by this version of the Tool.
//
//These parameters define the inputs needed for the Tool’s execution, including the expected data type and description for each input field. Structured as a stringified JSON schema, this format ensures the tool receives data in the expected format.
//
//version_description
//string
//Optional
//An optional description of the Tool version.
//
//fallback_content
//string
//Optional
//Optional text passed to the supplemental LLM in place of the tool call result. The LLM then uses this text to generate a response back to the user, ensuring continuity in the conversation if the Tool errors.
//
//description
//string
//Optional
//An optional description of what the Tool does, used by the supplemental LLM to choose when and how to call the function.

func (h *HumeApi) EviGetToolVersions(id string, pageNumber int64, pageSize int64, restrictToMostRecent bool) (*ToolsList, error) {
	h.printDebug("EviGetToolVersions")

	var toolsList ToolsList
	var err error

	if pageNumber < 0 {
		pageNumber = 0
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	toolsList.PageNumber = pageNumber
	toolsList.PageSize = pageSize

	url := fmt.Sprintf("%s/evi/tools/%s?page_number=%d&page_size=%d&restrict_to_most_recent=%t", h.BaseUrl, id, pageNumber, pageSize, restrictToMostRecent)
	h.printDebug(url)

	err = h.get(url, &toolsList)
	if err != nil {
		return nil, err
	}

	return &toolsList, nil
}
