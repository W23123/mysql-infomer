package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// S3StorageProvider represents an S3 compatible bucket for storing Backups.
type S3StorageProvider struct {
	// Region in which the S3 compatible bucket is located.
	Region string `json:"region"`
	// Endpoint (hostname only or fully qualified URI) of S3 compatible
	// storage service.
	Endpoint string `json:"endpoint"`
	// Bucket in which to store the Backup.
	Bucket string `json:"bucket"`
	// ForcePathStyle when set to true forces the request to use path-style
	// addressing, i.e., `http://s3.amazonaws.com/BUCKET/KEY`. By default,
	// the S3 client will use virtual hosted bucket addressing when possible
	// (`http://BUCKET.s3.amazonaws.com/KEY`).
	ForcePathStyle bool `json:"forcePathStyle"`
	// CredentialsSecret is a reference to the Secret containing the
	// credentials authenticating with the S3 compatible storage service.

	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

// StorageProvider defines the configuration for storing a Backup in a storage
// service.
type StorageProvider struct {
	S3 *S3StorageProvider `json:"s3"`
}

// MysqlBackupSpec defines the desired state of MysqlBackup
type MysqlBackupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	// StorageProvider configures where and how backups should be stored.
	StorageProvider StorageProvider `json:"storageProvider"`
	// Cluster is the Cluster to backup.
	Cluster *corev1.LocalObjectReference `json:"cluster"`
	// Backup name
	Backup *corev1.LocalObjectReference `json:"backup"`
}

const (
	BackupBegin   string = "Begin"
	BackupRunning        = "Running"
	BackupSuccess        = "Success"
	BackupFailed         = "Failed"
)

// MysqlBackupStatus defines the observed state of MysqlBackup
type MysqlBackupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
	State string `json:"state"`

	Location string `json:"location"`
	// TimeStarted is the time at which the backup was started.
	// +optional
	TimeStarted metav1.Time `json:"timeStarted"`
	// TimeCompleted is the time at which the backup completed.
	// +optional
	TimeCompleted metav1.Time `json:"timeCompleted"`
	Reason        string      `json:"reason,omitempty"`
}

// MysqlBackup is the Schema for the mysqlbackups API
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resourceName=mysqlbackups
type MysqlBackup struct {
	metav1.TypeMeta   `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +optional
	Spec   MysqlBackupSpec   `json:"spec,omitempty"`
	// +optional
	Status MysqlBackupStatus `json:"status,omitempty"`
}


// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// MysqlBackupList contains a list of MysqlBackup
type MysqlBackupList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MysqlBackup `json:"items"`
}
