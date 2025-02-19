#!/bin/bash

OUTPUT_BINARY="GoWeb"

VERSION="0.1.0"
PRODUCT_NAME="GoWeb"
COMPANY_NAME="onihilist"
COPYRIGHT="© 2024-2025 onihilist. All Rights Reserved"

cat <<EOF > version.rc
VS_VERSION_INFO VERSIONINFO
FILEVERSION    0,1,0
PRODUCTVERSION 0,1,0
{
    BLOCK "StringFileInfo"
    {
        BLOCK "040904b0"
        {
            VALUE "CompanyName",        "$COMPANY_NAME\0"
            VALUE "FileDescription",    "A golang website\0"
            VALUE "FileVersion",        "$VERSION\0"
            VALUE "LegalCopyright",     "$COPYRIGHT\0"
            VALUE "OriginalFilename",   "$PRODUCT_NAME\0"
            VALUE "ProductName",        "$PRODUCT_NAME\0"
            VALUE "ProductVersion",     "$VERSION\0"
        }
    }
    BLOCK "VarFileInfo"
    {
        VALUE "Translation", 0x409, 1200
    }
}
EOF

rsrc -o version.syso version.rc

go build -o "$OUTPUT_BINARY"

rm version.rc version.syso

echo "Build completed: $OUTPUT_BINARY"