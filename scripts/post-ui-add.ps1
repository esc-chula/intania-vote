# Check if the target path is provided
if ($args.Length -eq 0) {
    Write-Host "Usage: script.ps1 <target-path>"
    exit 1
}

# Check if the target directory exists
$TARGET_DIR = $args[0]
if (-not (Test-Path -Path $TARGET_DIR -PathType Container)) {
    Write-Host "The specified directory $TARGET_DIR does not exist."
    exit 1
}

# Define the target index file
$SRC_INDEX = "$TARGET_DIR/src/index.ts"
$COMPONENTS_DIR = "$TARGET_DIR/src/components/ui"
$COMPONENTS_INDEX = "$COMPONENTS_DIR/index.ts"

# Clear the index.ts file, or create it if it doesn't exist
Clear-Content -Path $COMPONENTS_INDEX -Force

# Loop through the .ts and .tsx files in the directory
$files = Get-ChildItem -Path $COMPONENTS_DIR | Where-Object { $_.Extension -match '\.ts$|\.tsx$' }
foreach ($file in $files) {
    $FILENAME = $file.Name
    $BASENAME = [System.IO.Path]::GetFileNameWithoutExtension($FILENAME)

    if ($BASENAME -ne "index") {
        Add-Content -Path $COMPONENTS_INDEX -Value "export * from './$BASENAME';"

        (Get-Content -Path $file.FullName) | ForEach-Object {
            $_ -replace '~/lib/utils', '../../lib/utils' -replace '~/components/ui', './'
        } | Set-Content -Path $file.FullName
    }
}

Write-Host "Updated successfully."
